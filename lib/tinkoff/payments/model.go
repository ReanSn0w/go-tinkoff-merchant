package payments

// MARK: - Init
// Ниже описаны структуры для обработки запроса на создание платежа

type InitRequest struct {
	TerminalKey     string // Идентиифкатр терминала выдаваемый банком
	Amount          int64  // Сумма в копейках
	OrderId         string // Идентификатор заказа в системе продавца
	Token           string `json:"Token,omitempty"`           // Подпись запроса
	IP              string `json:"IP,omitempty"`              // IP пользователя
	Description     string `json:"Description,omitempty"`     // Краткое описание
	Currency        int64  `json:"Currency,omitempty"`        // Код валюты в формате ISO421
	CustomerKey     string `json:"CustomerKey"`               // Идентификатор пользователя в системе продавца, Обязателен для рекурентных платежей
	Recurrent       string `json:"Recurrent,omitempty"`       // [Y] метка рекуррентного платежа
	PayType         string `json:"PayType,omitempty"`         // [O, T] однофазный или двухфазный платеж
	Language        string `json:"Language,omitempty"`        // [ru, en] Язык для платежной формы банка
	NotificationURL string `json:"NotificationURL,omitempty"` // URL для нотификаций на сайте продавца
	SuccessURL      string `json:"SuccessURL,omitempty"`      // URL для перенаправления пользователя при успешном платеже
	FailURL         string `json:"FailURL,omitempty"`         // URL для перенаправления пользователя при ошибке в платеже
	RedirectDueDate string `json:"RedirectDueDate,omitempty"` // [2016-08-31T12:28:00+03:00] Дата окончания действия ссылки
	Data            struct {
		Device        string `json:"Device,omitempty"`        // [SDK, Desctop, Mobile]
		DeviceOs      string `json:"DeviceOs"`                // ОС устройства
		DeviceWebView bool   `json:"DeviceWebView,omitempty"` // Признак запуска в WebView
		DeviceBrowser string `json:"Browser,omitempty"`       // Название браузера
		TinkoffPayWeb bool   `json:"TinkoffPayWeb,omitempty"` // Признак проведения платежа через TinkoffPay
	} `json:"DATA,omitempty"` // Дополнительные параметры платежа
	Receipt     Receipt `json:"Receipt,omitempty"`    // Данные чека
	Shops       []Shop  `json:"Shops,omitempty"`      // Объект с данными для распределения платежа по магазинам
	Descriptior string  `json:"Descriptor,omitempty"` // Динамический дескриптор точки
}

// При ffd 1.05
type Receipt struct {
	Items      []Item   // Массив товаров в чеке
	FfdVersion string   `json:"FfdVersion,omitempty"` // [1.2, 1.05]
	Email      string   `json:"Email,omitempty"`      // Email пользователя
	Phone      string   `json:"Phone,omitempty"`      // Телефон пользователя
	Taxation   string   // [osn, usn_income, envd, esn, patent]
	Payments   Payments `json:"Payemnts,omitempty"` // Объект c информацией о видах суммы платежа
}

type Shop struct {
	ShopCode string // Код продавца присвоенный банком
	Amount   int64  // Сумма в копейках для перевода продавцу
	Name     string `json:"Name,omitempty"` // Наименование позиции
	Fee      string `json:"Fee,omitempty"`  // Сумма комиссии (если не передано используется значение указанное при регистрации)
}

type Item struct {
	Name     string // Наименование товара
	Price    int64  // Цена в копейках
	Quantity int64  // Колличество
	Amount   int64  // Цена в копейках

	// Признак способа рассчета
	// [full_prepayment, prepayment, advance, full_payment, partial_payment, credit, credit_payment]
	PaymentMethod string `json:"PaymentMethod,omitempty"`

	// Признак предмета рассчета
	// [commodity, excise, job, service, gambling_bet, gambling_prize,
	//  lottery, lottery_prize, intellectual_activity, payment, agent_commission,
	//  composite, another]
	PaymentObject string `json:"PaymentObject,omitempty"`

	// Ставка налога
	// [none, vat0, vat10, vat20, vat110, vat120]
	Tax string

	Ean13    string `json:"Ean13,omitempty"`    // Штрих код в форамате, который требует касса
	ShopCode string `json:"ShopCode,omitempty"` // Код магазина

	// AgentData struct {
	// } `json:"AgentData,omitempty"` // Агентсткая схема. Параметробязателен при спользовании агентской схемы

	// SupplierInfo struct {
	// } `json:"SupplierInfo,omitempty"` // Данные поставщика платежного агента
}

type Payments struct {
	Electronic    int64 // Безналичные
	Cash          int64 `json:"Cash,omitempty"`          // Наличые
	AdvacePayment int64 `json:"Advicepayment,omitempty"` // Предварительная отплата
	Credit        int64 `json:"Credit,omitempty"`        // Постоплата (Кредит)
	Provision     int64 `json:"Provision,omitempty"`     // Иная форма оплаты
}

type InitResponse struct {
	TerminalKey string // Идентификатор терминала
	Amount      int64  // Суммма в копейках
	OrderId     string // Идентификатор платежа в системе продавца
	Success     bool   // Признак успешности опреации
	Status      string // Статус транзакции
	PaymentId   string // Уникальный идентификатор тразакции в системе банка
	ErrorCode   string // 0 - если успешно
	Message     string `json:"Message,omitempty"` // Сообщение об ошибке
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
	PaymentURL  string `json:"PaymentURL"`        // Ссылка на платежную форму
}

// MARK: - Confirm
// Ниже приведены структуры для запроса о подтверждении списания средств

type ConfirmRequest struct {
	TerminalKey string  // Идентификатор терминала выдаваемый банком
	PaymentId   string  // Идентификатор транзакции выдаваемый банком
	Token       string  // Подпись запроса
	IP          string  `json:"IP,omitempty"`      // IP пользователя
	Amount      int64   `json:"Amount,omitempty"`  // Сумма в копейках
	Reciept     Receipt `json:"Reciept,omitempty"` // Данные чека
	Shops       []Shop  `json:"Shops,omitempty"`   // Данные маркетплейсов
	Route       string  `json:"Route,omitempty"`   // [ТСВ, BNPL] Способ платежа
	Source      string  `json:"Source,omitempty"`  // [Installment, BNPL] Источник платежа
}

type ConfirmResponse struct {
	TerminalKey string // Идентификатор терминала
	Amount      int64  // Суммма в копейках
	OrderId     string // Идентификатор платежа в системе продавца
	Success     bool   // Признак успешности опреации
	Status      string // Статус транзакции
	PaymentId   string // Уникальный идентификатор тразакции в системе банка
	ErrorCode   string // 0 - если успешно
	Message     string `json:"Message,omitempty"` // Сообщение об ошибке
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
	Params      struct {
		Route        string `json:"Route,omitempty"`        // [TCB] способ платежа
		Source       string `json:"Source,omitempty"`       // [Installment] источник платежа
		CredotAmount int64  `json:"CreditAmount,omitempty"` // Сумма выданного платежа в копейках
	} `json:"Params,omitempty"` // Для платежей в рассрочку
}

// MARK: - Charge
// Ниже приведены структуды для выполнения рекуррентного платежа

type ChargeRequest struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	PaymentId   string // Идентификатор транзации выданные при метода Init
	RebillId    string // Идентификатор рекррентного пллатежа
	Token       string // Подпись запроса
	IP          string `json:"IP,omitempty"`        // IP клинта
	SendMail    bool   `json:"SendMail,omitempty"`  // Отправить уведомление покупателю на почту
	InfoEmail   string `json:"InfoEmail,omitempty"` // Email пользователя на который следует отправить уведомление
}

type ChargeResponse struct {
	TerminalKey string // Идентификатор терминала
	Amount      int64  // Суммма в копейках
	OrderId     string // Идентификатор платежа в системе продавца
	Success     bool   // Признак успешности опреации
	Status      string // Статус транзакции
	PaymentId   string // Уникальный идентификатор тразакции в системе банка
	ErrorCode   string // 0 - если успешно
	Message     string `json:"Message,omitempty"` // Сообщение об ошибке
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
}

// MARK: - Cancel
// Ниже приведены структуры для выполения честичной или полной отмены платежа

type CancelRequest struct {
	TerminalKey       string  // Идентификатор терминала выдаваемый банком
	PaymentId         string  // Идентификатор транзации выданные при метода Init
	Token             string  // Подпись запроса
	IP                string  `json:"IP,omitempty"`               // IP клинта
	Amount            int64   `json:"Amount,omitempty"`           // Суммма в копейках
	Reciept           Receipt `json:"Reciept,omitempty"`          // Список товаров по которым производится отмена платежа
	Shops             []Shop  `json:"Shops,omitempty"`            // JSON объект с данными маркетплейса
	Route             string  `json:"Route,omitempty"`            // [ТСВ, BNPL] Способ платежа
	Source            string  `json:"Source,omitempty"`           // [Installment, BNPL] Источник платежа
	QrMemberId        string  `json:"QrMemberId,omitempty"`       // Код банка в классификации СБП
	ExternalRequestId string  `json:"ExtenalRequestId,omitempty"` // Идентификатор транзакции на стороне мерчанта
}

type CancelResponse struct {
	TerminalKey       string // Идентификатор терминала
	Success           bool   // Признак успешности опреации
	Status            string // Статус транзакции
	PaymentId         string // Уникальный идентификатор тразакции в системе банка
	ErrorCode         string // 0 - если успешно
	OrderId           string // Идентификатор заказа в системе продавца
	OriginalAmount    int64  // Сумма в копейках до отмены
	NewAmount         int64  // Сумма операции после отмены
	Message           string `json:"Message,omitempty"`          // Сообщение об ошибке
	Details           string `json:"Details,omitempty"`          // Подробное описание ошибки
	ExternalRequestId string `json:"ExtenalRequestId,omitempty"` // Идентификатор транзакции на стороне мерчанта
}

// MARK: - GetState
// Ниже приведены структуры для выполнения запроса по получению статуса платежа

type GetStateRequest struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	PaymentId   string // Идентификатор транзации выданные при метода Init
	Token       string // Подпись запроса
	IP          string `json:"IP,omitempty"` // IP клинта
}

type GetStateResponse struct {
	TerminalKey string // Идентификатор терминала
	Amount      int64  // Суммма в копейках
	OrderId     string // Идентификатор платежа в системе продавца
	Success     bool   // Признак успешности опреации
	Status      string // Статус транзакции
	PaymentId   string // Уникальный идентификатор тразакции в системе банка
	ErrorCode   string // 0 - если успешно
	Message     string `json:"Message,omitempty"` // Сообщение об ошибке
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
	// Params      struct {
	// 	Route        string `json:"Route,omitempty"`        // [TCB] способ платежа
	// 	Source       string `json:"Source,omitempty"`       // [Installment] источник платежа
	// 	CredotAmount int64  `json:"CreditAmount,omitempty"` // Сумма выданного платежа в копейках
	// } `json:"Params,omitempty"` // Для платежей в рассрочку
}

// MARK: - CheckOrder
// Ниже приведены структуры для проверки статуса заказа

type CheckOrderRequest struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	OrderId     string // Номер заказа в стистеме продавца
	Token       string // Подпись запроса
}

type СheckOrderResponse struct {
	TerminalKey string               // Идентификатор терминала
	OrderId     string               // Идентификатор платежа в системе продавца
	Payments    []CheckOrderPayments // Детали
	Success     bool                 // Признак успешности опреации
	ErrorCode   string               // 0 - если успешно
	Message     string               `json:"Message,omitempty"` // Сообщение об ошибке
	Details     string               `json:"Details,omitempty"` // Подробное описание ошибки
}

type CheckOrderPayments struct {
	PaymentID string // Уникальный идентификатор транзакиции
	Amount    int64  `json:"Amount,omitempty"` // Суммма в копейках
	Status    string // Статус транзакции
	RRN       string `json:"RRN,omitempty"` // RRN опреции
	Success   string // Успешность прохождения запроса
	ErrorCode string // 0 - если успешно
	Message   string `json:"Message,omitempty"` // Сообщение об ошибке
}

// MARK: - SendClosingReciept
// Ниже описаны структуры запроса и ответа для отправки закрывающего чека

type SendClosingReceiptRequest struct {
	TerminalKey string  // Идентификатор терминала выдаваемый банком
	PaymentID   string  // Номер заказа в стистеме продавца
	Token       string  // Подпись запроса
	Receipt     Receipt // Массив данных чека
}

type SendClosingReceiptResponse struct {
	Status    string // Статус транзакции
	ErrorCode string // 0 - если успешно
	Message   string `json:"Message,omitempty"` // Сообщение об ошибке
}
