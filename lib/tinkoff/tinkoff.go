package tinkoff

import (
	"encoding/json"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
)

type Tinkoff struct {
	debug bool
	log   utils.Logger
	cl    *http.Client
}

func (t *Tinkoff) Debug() bool {
	return t.debug
}

func (t *Tinkoff) Logger() utils.Logger {
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
