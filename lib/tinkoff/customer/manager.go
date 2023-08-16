package customer

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/notifications"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

var (
	customerURLTest = "https://rest-api-test.tinkoff.ru/v2"
	customerURL     = "https://securepay.tinkoff.ru/v2"
)

func New(t utils.TinkoffService, terminalKey, password string) *Manager {
	return &Manager{
		service:     t,
		terminalKey: terminalKey,
		password:    password,
	}
}

type Manager struct {
	service     utils.TinkoffService
	terminalKey string
	password    string
	customURL   string
}

func (m *Manager) SetCustomURL(url string) {
	m.customURL = url
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

// Remove удаляет пользователя
func (m *Manager) Remove(customerKey string) (*RemoveResponse, error) {
	data := m.buildRequest(customerKey)
	result := &RemoveResponse{}
	err := m.request("/RemoveCustomer", http.MethodPost, data, result)
	return result, err
}

// ListCards получает список карт пользователя
func (m *Manager) ListCards(customerKey string) ([]CardItem, error) {
	data := m.buildRequest(customerKey)
	result := []CardItem{}
	err := m.request("/GetCardList", http.MethodPost, data, &result)
	return result, err
}

// AddCard Привязывает карту пользователю
func (m *Manager) AddCard(customerKey string, mods ...RequestModificator) (*AddCardResponse, error) {
	data := m.buildRequest(customerKey, mods...)
	result := &AddCardResponse{}
	err := m.request("/AddCard", http.MethodPost, data, result)
	return result, err
}

func (m *Manager) RemoveCard(customerKey, cardID string) (*RemoveCardResponse, error) {
	data := m.buildRequest(customerKey, WithCardID(cardID))
	result := &RemoveCardResponse{}
	err := m.request("/RemoveCard", http.MethodPost, data, result)
	return result, err
}

// Hnadler - создает handler для получения уведомлений
func (p *Manager) Handler(action func(item notifications.CardItem) error) func(http.ResponseWriter, *http.Request) {
	return notifications.
		New(p.service.Log(), p.terminalKey, p.password).
		Card(action)
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

	sign := signature.MakeSignature(r, p.password)
	r.Token = sign
	return r
}

func (p *Manager) buildURL(path string) string {
	if p.customURL != "" {
		return p.customURL + path
	}

	if p.service.Debug() {
		return customerURLTest + path
	}

	return customerURL + path
}
