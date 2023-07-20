package partner

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/token"
)

const (
	partnerRegistrationURLTest = "https://sm-register-test.tcsbank.ru"
	partnerRegistrationURL     = "https://sm-register.tinkoff.ru"
)

// New создает структуру для управления партнерами в банке
func New(service utils.TinkoffService, terminal string, password string) (*Partner, error) {
	t, err := token.New(service.Log(), partnerURL(service.Debug())+"/oauth/token", terminal, password)
	if err != nil {
		return nil, err
	}

	return &Partner{
		service: service,
		token:   t,
	}, nil
}

type Partner struct {
	service utils.TinkoffService
	token   *token.Token
}

// Register - метод для регистрации нового партнера в банке
func (service *Partner) Register(form RegistrationRequest) (*Response, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(form)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, partnerURL(service.service.Debug())+"/register", buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+service.token.Get())

	response := &Response{}
	err = service.service.Request(request, response)
	return response, err
}

// Update - метод для обновления информации о партнере
func (service *Partner) Update(shopID string, account BankAccount) (*Response, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(map[string]interface{}{
		"bankAccount": account,
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPatch, partnerURL(service.service.Debug())+"/register/"+shopID, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+service.token.Get())

	response := &Response{}
	err = service.service.Request(request, response)
	return response, err
}

func partnerURL(debug bool) string {
	if debug {
		return partnerRegistrationURLTest
	}

	return partnerRegistrationURL
}
