package notifications

type Item struct {
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
