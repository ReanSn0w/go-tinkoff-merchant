package customer

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

var (
	customerURLTest = "https://rest-api-test.tinkoff.ru/v2"
	customerURL     = "https://securepay.tinkoff.ru/v2"
)

func New(t utils.TinkoffService, terminalKey string) *Manager {
	return &Manager{
		service:     t,
		terminalKey: terminalKey,
	}
}

type Manager struct {
	service     utils.TinkoffService
	terminalKey string
}

// Add добавляет нового пользователя
func (m *Manager) Add(customerKey string, mods ...RequestModificator) (*AddResponse, error) {
	data := m.buildRequest(customerKey, mods...)
	result := &AddResponse{}
	err := m.request("/AddCustomer", http.MethodPost, data, result)
	return result, err
}

// Get получет инофрмацию о пользователе
func (m *Manager) Get(customerKey string) (*GetResponse, error) {
	data := m.buildRequest(customerKey)
	result := &GetResponse{}
	err := m.request("/GetCustomer", http.MethodPost, data, result)
	return result, err
}

func (m *Manager) Remove(customerKey string) (*RemoveResponse, error) {
	data := m.buildRequest(customerKey)
	result := &RemoveResponse{}
	err := m.request("/RemoveCustomer", http.MethodPost, data, result)
	return result, err
}

func (m *Manager) ListCards(customerKey string) ([]CardItem, error) {
	data := m.buildRequest(customerKey)
	result := []CardItem{}
	err := m.request("/GetCardList", http.MethodPost, data, result)
	return result, err
}

func (p *Manager) request(path string, method string, request, response interface{}) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, p.buildURL(path), buffer)
	if err != nil {
		return err
	}

	return p.service.Request(req, response)
}

func (p *Manager) buildRequest(customerKey string, mods ...RequestModificator) Request {
	r := Request{
		TerminalKey: p.terminalKey,
		CustomerKey: customerKey,
	}

	for index := range mods {
		r = mods[index](r)
	}

	sign := signature.MakeSignature(r)
	r.Token = sign
	return r
}

func (p *Manager) buildURL(path string) string {
	if p.service.Debug() {
		return customerURLTest + path
	}

	return customerURL + path
}
