package payments

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/notifications"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

const (
	paymentsURLTest = "https://rest-api-test.tinkoff.ru/v2"
	paymentsURL     = "https://securepay.tinkoff.ru/v2"
)

// New создает новую структуру для работы с платежами
func New(service utils.TinkoffService, terminalID, password string) (*Manager, error) {
	return &Manager{
		service:    service,
		terminalID: terminalID,
		password:   password,
	}, nil
}

// PaymentsManager содержит методы для работы с платежами
type Manager struct {
	service    utils.TinkoffService
	terminalID string
	password   string
}

// Hnadler - создает handler для получения уведомлений
func (p *Manager) Handler(action func(item *notifications.Item) error) func(http.ResponseWriter, *http.Request) {
	return notifications.
		New(p.service.Log(), p.terminalID, p.password).
		HandlerFunc(action)
}

// Init инициирует платежную сессию
func (p *Manager) Init(data InitRequest) (*InitResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &InitResponse{}
	err := p.request("/Init", http.MethodPost, data, result)
	return result, err
}

// Confirm Осуществляет списание заблокированных денежных средств
func (p *Manager) Confirm(data ConfirmRequest) (*ConfirmResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &ConfirmResponse{}
	err := p.request("/Confirm", http.MethodPost, data, result)
	return result, err
}

// Charge осуществляет списание рекуррентного платежа
func (p *Manager) Charge(data ChargeRequest) (*ChargeResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &ChargeResponse{}
	err := p.request("/Charge", http.MethodPost, data, result)
	return result, err
}

// Cancel метод для отмены плетежа
func (p *Manager) Cancel(data CancelRequest) (*CancelResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &CancelResponse{}
	err := p.request("/Cancel", http.MethodPost, data, result)
	return result, err
}

// GetState возвращает статус платежа
func (p *Manager) GetState(data GetStateRequest) (*GetStateResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &GetStateResponse{}
	err := p.request("/GetState", http.MethodPost, data, result)
	return result, err
}

func (p *Manager) CheckOrder(data CheckOrderRequest) (*СheckOrderResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &СheckOrderResponse{}
	err := p.request("/CheckOrder", http.MethodPost, data, result)
	return result, err
}

func (p *Manager) SendClosingReceipt(data SendClosingReceiptRequest) (*SendClosingReceiptResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &SendClosingReceiptResponse{}
	err := p.request("/GetState", http.MethodPost, data, result)
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

func (p *Manager) buildURL(path string) string {
	if p.service.Debug() {
		return paymentsURLTest + path
	}

	return paymentsURL + path
}
