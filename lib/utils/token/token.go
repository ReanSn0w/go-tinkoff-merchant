package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
)

func New(logger utils.Logger, baseURL, login, password string) (*Token, error) {
	token := &Token{
		baseURL:  baseURL,
		login:    login,
		password: password,
		mutex:    &sync.RWMutex{},
		done:     make(chan int),
		logger:   logger,
		cl: &http.Client{
			Timeout: time.Second * 3,
		},
	}

	if err := token.refreshOnce(); err != nil {
		return nil, err
	}

	go token.refreshTask()
	return token, nil
}

type Token struct {
	baseURL  string
	login    string
	password string
	token    *response

	logger utils.Logger
	cl     *http.Client
	mutex  *sync.RWMutex
	timer  *time.Timer
	done   chan int
}

type response struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

func (t *Token) Get() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.token == nil {
		return ""
	}

	return t.token.AccessToken
}

func (t *Token) Stop() {
	t.done <- 0
}

func (t *Token) refreshTask() {
	t.timer = time.NewTimer(time.Second * 60)
	restart := 0

	for {
		if t.timer == nil {
			break
		}

		if restart == 0 {
			t.timer.Reset(time.Hour)
		} else {
			t.timer.Reset(time.Duration(restart) * time.Second)
			restart = 0
		}

		select {
		case <-t.timer.C:
			if err := t.refreshOnce(); err != nil {
				restart = 10
				t.logger.Logf("[WARN] error with update token. next update in 10 sec")
			}
		case <-t.done:
			t.timer.Stop()
			t.timer = nil
		}
	}
}

func (t *Token) refreshOnce() error {
	body := t.makeBody()
	req, err := http.NewRequest("POST", t.baseURL, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth("partner", "partner")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.cl.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(resp.Body)
		t.logger.Logf("[WARN] body: %v", buffer.String())
		return errors.New("response code != 200")
	}

	tokens := &response{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(tokens)
	if err != nil {
		return err
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.token = tokens
	return nil
}

func (t *Token) makeBody() io.Reader {

	return bytes.NewReader([]byte(fmt.Sprintf("grant_type=password&username=%v&password=%v", t.login, t.password)))

	// buffer := new(bytes.Buffer)

	// mpf := multipart.NewWriter(buffer)
	// _ = mpf.WriteField("grant_type", "password")
	// _ = mpf.WriteField("username", t.login)
	// _ = mpf.WriteField("password", t.password)

	// return buffer
}
