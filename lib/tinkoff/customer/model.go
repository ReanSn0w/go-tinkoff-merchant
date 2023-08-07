package customer

type Request struct {
	TerminalKey string // Идентификатор терминала выданый банком
	CustomerKey string // Идентификатор пользователя
	Token       string // Подпись запроса

	// Метод проверки карты
	// Default: NO
	// Case: NO - сохранение карты без проверки
	// Case: HOLD - списание на 0 рублей
	// Case: 3DS - проверка по протоколу 3-D Secure
	// Case: 3DSHOLD - проверка по протоколу 3-D Secure и списание на 0 рублей
	CheckType string `json:"CheckType,omitempty"`

	CardId string `json:"CardId,omitempty"` // Идентификатор карты

	PayForm       string `json:"PayForm,omitempty"`       // Название платежной формы
	ResidentState *bool  `json:"ResidentState,omitempty"` // Признак резидента РФ
	IP            string `json:"IP,omitempty"`            // IP адрес пользователя
	Email         string `json:"Email,omitempty"`         // Email клиента
	Phone         string `json:"Phone,omitempty"`         // Телефон клиента
}

// MARK: - AddCustomer
// Ниже описаны структуры для добавления пользователей, для последующей привязки карты

type AddResponse struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	CustomerKey string // Идентификатор пользователя
	Success     bool   // Успешность прохождения запроса
	ErrorCode   string // Код ошибки, 0 - для успешной операции
	Message     string `json:"Message,omitempty"` // Текст сообщения
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
}

// MARK: - GetCustomer
// Ниже приведены структуры для описания ответа/запроса на получение информации о пользователе

type GetResponse struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	CustomerKey string // Идентификатор пользователя
	Success     bool   // Успешность прохождения запроса
	ErrorCode   string // Код ошибки, 0 - для успешной операции
	Message     string `json:"Message,omitempty"` // Текст сообщения
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
	Email       string `json:"Email,omitempty"`   // Email клиента
	Phone       string `json:"Phone,omitempty"`   // Телефон клиента
}

// MARK: - RemoveCustomer
// Ниже приведены структуры для ответа/запроса на удаление пользователя

type RemoveResponse struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	CustomerKey string // Идентификатор пользователя
	Success     bool   // Успешность прохождения запроса
	ErrorCode   string // Код ошибки, 0 - для успешной операции
	Message     string `json:"Message,omitempty"` // Текст сообщения
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
}

// MARK: - GetCardList
// Ниже приведены структуры для ответа запроса на полечение списка привязанных карт у пользователя

type CardItem struct {
	CardId   string // Идентификатор карты в системе банка
	Pan      string // Номер карты со скрытой частью
	Status   string // [A, I, D] Статус карты
	RebillId string `json:"RebillId,omitempty"` // Идентификатор рекуррентного платежа
	CardType string `json:"CardType,omitempty"` // [0,1,2] (читай как списание, пополнение, списание и пополнение) Тип карты
	ExpDate  string `json:"ExpDate,omitempty"`  // Срок действия карты
}

type AddCardResponse struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	CustomerKey string // Идентификатор пользователя
	RequestKey  string // Идентификатор запроса
	PaymentURL  string // URL для перехода на страницу проверки карты
	Success     bool   // Успешность прохождения запроса
	ErrorCode   string // Код ошибки, 0 - для успешной операции
	Message     string `json:"Message,omitempty"` // Текст сообщения
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
}

type RemoveCardResponse struct {
	TerminalKey string // Идентификатор терминала выдаваемый банком
	CustomerKey string // Идентификатор пользователя
	CardId      string // Идентификатор карты в системе банка
	Success     bool   // Успешность прохождения запроса
	ErrorCode   string // Код ошибки, 0 - для успешной операции
	Message     string `json:"Message,omitempty"` // Текст сообщения
	Details     string `json:"Details,omitempty"` // Подробное описание ошибки
}
