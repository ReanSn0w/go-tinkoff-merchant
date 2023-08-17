package safedeal_test

import (
	"testing"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff"
	"github.com/go-pkgz/lgr"
)

func TestSafeDeal_AddCard(t *testing.T) {
	sd, err := tinkoff.New(false, lgr.New(lgr.Debug)).SafeDeal(
		"1683019138816",
		"pf1h19bhd7ozster",
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := sd.Customer().AddCard("reansnow_001")

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("%v", resp.PaymentURL)
}
