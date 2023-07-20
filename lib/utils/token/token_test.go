package token_test

import (
	"testing"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/token"
	"github.com/go-pkgz/lgr"
)

func Test_New(t *testing.T) {
	tok, err := token.New(lgr.Default(), "https://sm-register-test.tcsbank.ru", "aceplace", "aceplace")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("token: ", tok.Get())
}
