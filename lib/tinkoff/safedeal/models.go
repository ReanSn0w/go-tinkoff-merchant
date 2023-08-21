package safedeal

type InitPayout struct {
	TerminalKey string
	Token       string
	OrderId     string
	CardId      string
	Amount      int64
	CustomerKey string
	Data        struct {
		SpAccumulationId string `json:"SpAccumulationId,omitempty"`
	} `json:"DATA,omitempty"`
}

type InitResponse struct {
	TerminalKey string
	Amount      int64
	OrderId     string
	Success     bool
	Status      string
	PaymentId   string
	ErrorCode   string // id 0 - то без ошибок
	Mesasge     string `json:"Message,omitempty"`
	Details     string `json:"Details,omitempty"`
}

type PaymentRequest struct {
	TerminalKey string
	PaymentId   string
	Token       string
}

type PaymentResponse struct {
	TerminalKey string
	OrderId     string
	Success     bool
	Status      string
	PaymentId   string
	ErrorCode   string // id 0 - то без ошибок
	Mesasge     string `json:"Message,omitempty"`
	Details     string `json:"Details,omitempty"`
}

type StateResponse struct {
	TerminalKey string
	OrderId     string
	Success     bool
	Status      string
	PaymentId   string
	ErrorCode   string // id 0 - то без ошибок
	Mesasge     string `json:"Message,omitempty"`
	Details     string `json:"Details,omitempty"`
}
