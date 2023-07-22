package tinkoff

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/customer"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/notifications"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/partner"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/payments"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
)

func New(debug bool, log utils.Logger) *Tinkoff {
	return &Tinkoff{
		debug: debug,
		log:   log,
		cl: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

type Tinkoff struct {
	debug bool
	log   utils.Logger
	cl    *http.Client
}

func (t *Tinkoff) Partner(username, password string) (*partner.Partner, error) {
	return partner.New(t, username, password)
}

func (t *Tinkoff) Payments(terminalID string) (*payments.PaymentManager, error) {
	return payments.New(t, terminalID)
}

func (t *Tinkoff) Notifications(terminalID string) *notifications.Manager {
	return notifications.New(t.log, terminalID)
}

func (t *Tinkoff) Customer(terminalID string) *customer.Manager {
	return customer.New(t, terminalID)
}

func (t *Tinkoff) Debug() bool {
	return t.debug
}

func (t *Tinkoff) Log() utils.Logger {
	return t.log
}

func (t *Tinkoff) Request(r *http.Request, data any) error {
	resp, err := t.cl.Do(r)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		err := decoder.Decode(data)
		return err
	default:
		errorStruct := &ErrorResponse{}
		err := decoder.Decode(errorStruct)
		if err != nil {
			return err
		}

		return errorStruct
	}
}
