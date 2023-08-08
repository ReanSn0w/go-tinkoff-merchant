package safedeal

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/customer"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/payments"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

const (
	safeDealURL     = "https://securepay.tinkoff.ru/e2c/v2"
	safeDealURLTest = "https://rest-api-test.tinkoff.ru/e2c/v2"
)

// New создает новую структуру для работы с платежами
func New(service utils.TinkoffService, terminalID, password string) (*Manager, error) {
	p, err := payments.New(service, terminalID, password)
	if err != nil {
		return nil, err
	}

	c := customer.New(service, terminalID, password)
	if service.Debug() {
		c.SetCustomURL(safeDealURLTest)
	} else {
		c.SetCustomURL(safeDealURL)
	}

	return &Manager{
		service:    service,
		terminalID: terminalID,
		password:   password,
		payments:   p,
		customer:   c,
	}, nil
}

// PaymentsManager содержит методы для работы с платежами
type Manager struct {
	service    utils.TinkoffService
	terminalID string
	password   string
	payments   *payments.Manager
	customer   *customer.Manager
}

func (m *Manager) Payment() *payments.Manager {
	return m.payments
}

func (m *Manager) Customer() *customer.Manager {
	return m.customer
}

func (m *Manager) InitPayout(req InitPayout) (*InitResponse, error) {
	req.TerminalKey = m.terminalID
	sign := signature.MakeSignature(req, m.password)
	req.Token = sign

	result := &InitResponse{}
	err := m.request("/Init", http.MethodPost, req, result)
	return result, err
}

func (m *Manager) PaymentPayout(paymentId string) (*PaymentResponse, error) {
	req := PaymentRequest{
		TerminalKey: m.terminalID,
		PaymentId:   paymentId,
	}
	sign := signature.MakeSignature(req, m.password)
	req.Token = sign

	result := &PaymentResponse{}
	err := m.request("/Payment", http.MethodPost, req, result)
	return result, err
}

func (m *Manager) GetStatePayout(paymentId string) (*StateResponse, error) {
	req := PaymentRequest{
		TerminalKey: m.terminalID,
		PaymentId:   paymentId,
	}
	sign := signature.MakeSignature(req, m.password)
	req.Token = sign

	result := &StateResponse{}
	err := m.request("/Payment", http.MethodPost, req, result)
	return result, err
}

func (m *Manager) request(path string, method string, request, response interface{}) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, m.buildURL(path), buffer)
	if err != nil {
		return err
	}

	return m.service.Request(req, response)
}

func (m *Manager) buildURL(path string) string {
	if m.service.Debug() {
		return safeDealURLTest + path
	}

	return safeDealURL + path
}
