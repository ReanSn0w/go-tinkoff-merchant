package payments

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

const (
	paymentsURLTest = "https://rest-api-test.tinkoff.ru/v2"
	paymentsURL     = "https://securepay.tinkoff.ru/v2"
)

// New создает новую структуру для работы с платежами
func New(service utils.TinkoffService, terminalID, password string) (*PaymentManager, error) {
	return &PaymentManager{
		service:    service,
		terminalID: terminalID,
		password:   password,
	}, nil
}

// PaymentsManager содержит методы для работы с платежами
type PaymentManager struct {
	service    utils.TinkoffService
	terminalID string
	password   string
}

// Init инициирует платежную сессию
func (p *PaymentManager) Init(data InitRequest) (*InitResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &InitResponse{}
	err := p.request("/Init", http.MethodPost, data, result)
	return result, err
}

// Confirm Осуществляет списание заблокированных денежных средств
func (p *PaymentManager) Confirm(data ConfirmRequest) (*ConfirmResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &ConfirmResponse{}
	err := p.request("/Confirm", http.MethodPost, data, result)
	return result, err
}

// Charge осуществляет списание рекуррентного платежа
func (p *PaymentManager) Charge(data ChargeRequest) (*ChargeResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &ChargeResponse{}
	err := p.request("/Charge", http.MethodPost, data, result)
	return result, err
}

// Cancel метод для отмены плетежа
func (p *PaymentManager) Cancel(data CancelRequest) (*CancelResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &CancelResponse{}
	err := p.request("/Cancel", http.MethodPost, data, result)
	return result, err
}

// GetState возвращает статус платежа
func (p *PaymentManager) GetState(data GetStateRequest) (*GetStateResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &GetStateResponse{}
	err := p.request("/GetState", http.MethodPost, data, result)
	return result, err
}

func (p *PaymentManager) CheckOrder(data CheckOrderRequest) (*СheckOrderResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &СheckOrderResponse{}
	err := p.request("/GetState", http.MethodPost, data, result)
	return result, err
}

func (p *PaymentManager) SendClosingReceipt(data SendClosingReceiptRequest) (*SendClosingReceiptResponse, error) {
	data.TerminalKey = p.terminalID
	sign := signature.MakeSignature(data, p.password)
	data.Token = sign

	result := &SendClosingReceiptResponse{}
	err := p.request("/GetState", http.MethodPost, data, result)
	return result, err
}

func (p *PaymentManager) request(path string, method string, request, response interface{}) error {
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

func (p *PaymentManager) buildURL(path string) string {
	if p.service.Debug() {
		return paymentsURLTest + path
	}

	return paymentsURL + path
}
