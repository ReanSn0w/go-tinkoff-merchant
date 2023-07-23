package notifications

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

func New(log utils.Logger, terminalKey, password string) *Manager {
	return &Manager{
		terminalKey: terminalKey,
	}
}

type Manager struct {
	terminalKey string
	password    string
	logger      utils.Logger
}

func (m *Manager) HandlerFunc(action func(*Item) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := m.checkIPAddress(r.RemoteAddr)
		if err != nil {
			m.log(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		item := &Item{}
		err = decoder.Decode(item)
		if err != nil {
			m.log(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sign := item.Token
		item.Token = ""
		if signature.MakeSignature(item, m.password) != sign {
			m.log(errors.New("body signature invalid"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if item.TerminalKey != m.terminalKey {
			m.log(errors.New("invald terminal key"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = action(item)
		if err != nil {
			m.log(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// производит проверку запроса на передачу уведомления по отношению к сетям Tinkoff
func (m *Manager) checkIPAddress(ip string) error {
	// TODO: - реализовать метод
	return nil
}

func (m *Manager) log(err error) {
	if m.logger != nil {
		m.logger.Logf("[WARN] notification error: %v", err)
	}
}
