package partner

// MARK: - Registration

type RegistrationRequest struct {
	ID                string      `json:"shopArticleId"`     // Идентификатор на нашей стороне
	BillingDescriptor string      `json:"billingDescriptor"` // Название в смс и 3s secure
	FullName          string      `json:"fullName"`          // Полное название организации
	Name              string      `json:"name"`              // Сокращенное название до 512 символов
	Inn               string      `json:"inn"`               // ИНН
	Kpp               string      `json:"kpp"`               // КПП передать 000000000 если нет
	Ogrn              int64       `json:"ogrn,omitempty"`    // ОГРН для компаний
	Smz               bool        `json:"smz,omitempty"`     // Флаг указывающий на самозанятых
	Addresses         []Address   `json:"addresses"`         // Адреса
	Email             string      `json:"email"`             // Email
	Ceo               Ceo         `json:"ceo"`               // Сведения о руководителе
	SiteURL           string      `json:"siteUrl"`           // ССылка на страницу AcePlace
	BankAccount       BankAccount `json:"bankAccount"`       // Банковский счет
}

type Address struct {
	Type        string  `json:"type"`    // legal - юридический actual - фактический post - почтовый other - прочий
	Zip         string  `json:"zip"`     // Индекс
	Country     string  `json:"country"` // Код страниы
	City        string  `json:"city"`    // Город
	Street      string  `json:"street"`  // Улица
	Description *string `json:"description,omitempty"`
}

type BankAccount struct {
	Account    string `json:"account"`    // Рассчетный счет
	KorAccount string `json:"korAccount"` // Корр счет
	BankName   string `json:"bankName"`   // Название банка
	Bik        string `json:"bik"`        // БИК
	Kbk        string `json:"kbk"`        // КБК
	Oktmo      string `json:"oktmo"`      // ОКТМО
	Details    string `json:"details"`    // Перевод средств по договору No 3333-3333 от 16.09.2021 по Реестру Операций от ${date}. Сумма комиссии ${rub} руб. ${kop} коп., НДС не облагается.
	Tax        int64  `json:"tax"`        // Процент отчислений в пользу маркетплейса
}

type Ceo struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName string  `json:"middleName"`
	BirthDate  string  `json:"birthDate"`
	Country    *string `json:"country,omitempty"`
}

type Fiscalization struct {
	Company   string `json:"company"`
	NotifyURL string `json:"notifyUrl"`
}

type Founders struct {
	Individuals []Ceo `json:"individuals"`
}

type License struct {
	Type        string `json:"type"`
	Number      string `json:"number"`
	IssueDate   string `json:"issueDate"`
	IssuedBy    string `json:"issuedBy"`
	ExpiryDate  string `json:"expiryDate"`
	Description string `json:"description"`
}

type Phone struct {
	Type        string `json:"type"` // common – основной fax – факс other - прочий
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type Response struct {
	ID        string        `json:"code"`
	ShopCode  int64         `json:"shopCode"`
	Terminals []interface{} `json:"terminals"`
}
