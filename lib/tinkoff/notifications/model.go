package notifications

type Item interface {
	GetTerminalKey() string
	GetToken() string
	RemoveToken() Item
}

type PaymentItem struct {
	TerminalKey  string                 // Идентификатор терминала выданный банком
	OrderID      string                 // Номер заказа в системеп продавци
	Success      bool                   // Успешность проведения операции
	Status       string                 // Статус платежа
	PaymentID    string                 // Идентификатор транзакции в системе банка
	ErrorCode    string                 // Код ошибки, 0 - если успешно
	Amount       int64                  // Сумма платежа в копейках
	CardId       int64                  // Идентификатор привязанной карты
	Pan          string                 // Зфмфскированный номер карты
	ExpDate      string                 // Срок действия карты
	RebillId     string                 // идентификатор для рекуррентного платежа
	Token        string                 // Подпись запроса
	DATA         map[string]interface{} // Дополнительные параметры запроса
	Route        string                 `json:"Route,omitempty"`        // Способ патежа
	Source       string                 `json:"Source"`                 // Источник платежа
	CreditAmount int64                  `json:"CreditAmount,omitempty"` // Сумма выданного кредита в копейках
}

func (p PaymentItem) GetTerminalKey() string {
	return p.TerminalKey
}

func (p PaymentItem) GetToken() string {
	return p.Token
}

func (p PaymentItem) RemoveToken() Item {
	p.Token = ""
	return p
}

type CardItem struct {
	TerminalKey string // Идентификатор терминала выданный банком
	Success     bool   // Успешность прохождения запроса
	Status      string // Статус платежа ["COMPLETED", "REJECTED"]
	Token       string // Подпись запроса

	CustomerKey      string `json:"CustomerKey,omitempty"`      // Идентификатор пользователя
	RequestKey       string `json:"RequestKey,omitempty"`       // Идентификатор запроса
	PaymentId        string `json:"PaymentId,omitempty"`        // Идентификатор транзакции в системе банка
	ErrorCode        string `json:"ErrorCode,omitempty"`        // Код ошибки, 0 - если успешно
	CardId           int64  `json:"CardId,omitempty"`           // Идентификатор карты в системе банка
	Pan              string `json:"Pan,omitempty"`              // Зашифрованный номер карты
	ExpDate          string `json:"ExpData,omitempty"`          // Срок действия карты
	RebillId         string `json:"RebillId,omitempty"`         // Идентификатор для рекуррентного платежа
	SpAccumulationId string `json:"SpAccumulationId,omitempty"` // Accumulation ID
}

func (c CardItem) GetTerminalKey() string {
	return c.TerminalKey
}

func (c CardItem) GetToken() string {
	return c.Token
}

func (c CardItem) RemoveToken() Item {
	c.Token = ""
	return c
}
