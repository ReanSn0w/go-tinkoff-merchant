package safedeal

// import "fmt"

// func NewSafeDeal(terminalKey, password string) *SafeDeal {
// 	return &SafeDeal{}
// }

// type SafeDeal struct {
// 	terminalKey string
// 	password    string
// }

// type InitRequest struct {
// 	TerminalKey      string           `json:"TerminalKey"`        // Идентификатор терминала
// 	OrderID          string           `json:"OrderId"`            // Идентификатор заказа в системе продавца
// 	IP               string           `json:"IP,omitempty"`       // IP адрес клиента
// 	CardID           string           `json:"CardId"`             // Идентификатор карты для пополнения
// 	Amount           int              `json:"Amount"`             // Сумма в копейках
// 	Currency         float64          `json:"Currency,omitempty"` // Идентификатор валюты ISO 4217
// 	CustomerKey      string           `json:"CustomerKey"`        // Идентификатор покупателя в системе продавца
// 	Data             *InitRequestData `json:"DATA,omitempty"`     // Дополнительные данные
// 	DigiestValue     string           `json:"DigestValue"`        // Хеш в base64
// 	SignatureValue   string           `json:"SignatureValue"`     // Подпись в base64
// 	X509SerianNumber string           `json:"X509SerianNumber"`   // Серийный номер сертификата
// }

// type InitRequestData struct {
// 	SLastname          string `json:"s_lastname"`          // Фамилия отправителя
// 	SFirstname         string `json:"s_firstname"`         // Имя отправителя
// 	SMiddlename        string `json:"s_middlename"`        // Отчество отправителя
// 	SBirthdate         string `json:"s_dateOfBirth"`       // Дата рождения в формате DD.MM.YYYY
// 	SCitizenship       string `json:"s_citizenhip"`        // Гражданство отправителя
// 	SPassportseries    string `json:"s_passportSeries"`    // Серия паспорта
// 	SPassportNumber    string `json:"s_passportNumber"`    // Номер паспорта
// 	SPassportIssueDate string `json:"s_passportIssueDate"` // Дата выдачи паспорта в формате DD.MM.YYYY
// 	SAccountNumber     string `json:"s_accountNumber"`     // Номер карты или рассчетного счета

// 	RLastname      string `json:"r_lastname"`       // Фамилия получателя
// 	RFirstname     string `json:"r_firstname"`      // Имя получателя
// 	RMiddlename    string `json:"r_middlename"`     // Отчество получателя
// 	AgrementNumber string `json:"agreement_number"` // Номер договора займа

// 	TDomestic string `json:"t_domestic"` // 0 - международный 1 - внутри страны
// }

// type InitResponse struct {
// 	Success   bool   `json:"Success"`           // Метка успешности
// 	ErrorCode string `json:"Error_code"`        // Код ошибки
// 	Message   string `json:"Message,omitempty"` // Сообщение об ошибке
// 	Details   string `json:"Details,omitempty"` // Детальное описание ошибки

// 	Status      string `json:"Status"`      // Checked, Rejected
// 	TerminalKey string `json:"TerminalKey"` // Идентификатор терминала
// 	Amount      int    `json:"Amount"`      // Сумма в копейках
// 	OrderID     string `json:"OrderId"`     // Идентификатор заказа в системе продавца
// 	PaymentID   string `json:"PaymentId"`   // Уникальный номер транзации в системе банка
// }

// func (ir *InitResponse) Error() error {
// 	if ir.ErrorCode == "0" {
// 		return nil
// 	}

// 	return fmt.Errorf("SafeDealError. \nCode: %s\nMessage: %s\nDetails: %s\n", ir.ErrorCode, ir.Message, ir.Details)
// }

// func (sd *SafeDeal) Init(req InitRequest) *InitResponse {

// }

// type PaymentResponse struct {
// 	Success   bool   `json:"Success"`           // Метка успешности
// 	ErrorCode string `json:"Error_code"`        // Код ошибки
// 	Message   string `json:"Message,omitempty"` // Сообщение об ошибке
// 	Details   string `json:"Details,omitempty"` // Детальное описание ошибки

// 	Status      string `json:"Status"`      // COMPLETED, REJECTED, CREDIT_CHECKING
// 	TerminalKey string `json:"TerminalKey"` // Идентификатор терминала
// 	OrderID     string `json:"OrderId"`     // Идентификатор заказа в системе продавца
// 	PaymentID   string `json:"PaymentId"`   // Уникальный номер транзации в системе банка
// }

// func (ir *PaymentResponse) Error() error {
// 	if ir.ErrorCode == "0" {
// 		return nil
// 	}

// 	return fmt.Errorf("SafeDealError. \nCode: %s\nMessage: %s\nDetails: %s\n", ir.ErrorCode, ir.Message, ir.Details)
// }

// func (sd *SafeDeal) Payment(paymentID string) *PaymentResponse {
// 	// {
// 	// 	"TerminalKey": " TinkoffBankTest ",
// 	// 	"PaymentId": "700000085140",
// 	// 	"DigestValue": "qfeohMmrsEvr4QPB8CeZETb+W6VDEGnMrf+oVjvSaMU=",
// 	// 	"SignatureValue": "rNTloWBbTsid1n9B1ANZ9/VasWJyg6jfiMeI12ERBSlOnzy6YFqMaa5nRb9ZrK9w
// 	// 	bKimIBD70v8j8eP/tKn7/g==",
// 	// 	"X509SerialNumber": "2613832945"
// 	// 	}
// }

// func (sd *SafeDeal) GetState(paymentID string) *GetStateResponse {

// }
