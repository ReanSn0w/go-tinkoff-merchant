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

func New(service utils.TinkoffService, terminal string, password string, debug bool) (*Partner, error) {
	t, err := token.New(service.Log(), partnerURL(debug), terminal, password)
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

func (service *Partner) Register(form RegistrationRequest) (*RegistrationResponse, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(form)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", partnerURL(service.service.Debug())+"/register", buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+service.token.Get())

	response := &RegistrationResponse{}
	err = service.service.Request(request, response)
	return response, err
}

func partnerURL(debug bool) string {
	if debug {
		return partnerRegistrationURLTest
	}

	return partnerRegistrationURL
}
