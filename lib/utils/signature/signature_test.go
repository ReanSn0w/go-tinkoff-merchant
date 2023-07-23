package signature_test

import (
	"testing"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

func Test_MakeSignature(t *testing.T) {
	pairs := []struct {
		Value    Data
		Password string
		Hash     string
	}{
		{
			Value: Data{
				Amount:      100000,
				Description: "test",
				OrderId:     "TokenExample",
				TerminalKey: "TinkoffBankTest",
			},
			Password: "TinkoffBankTest",
			Hash:     "48d4ca825aab2ede06736d3eae099bd56ac97bd1bcdd598aff210f729de4eb21",
		},
	}

	for _, item := range pairs {
		s := signature.MakeSignature(item.Value, item.Password)

		if s != item.Hash {
			t.Logf("signature not setted: \n value: %s \n hash: %s", s, item.Hash)
			t.Fail()
		}
	}
}

// MARK: - support

type Data struct {
	TerminalKey string
	Amount      int64
	Description string
	OrderId     string
}
