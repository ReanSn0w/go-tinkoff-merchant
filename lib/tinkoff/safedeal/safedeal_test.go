package safedeal_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/payments"
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

func TestSafeDeal_PaymentWSMID(t *testing.T) {
	sd, err := tinkoff.New(false, lgr.New(lgr.Debug)).SafeDeal(
		"1683019138816",
		"pf1h19bhd7ozster",
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := sd.Payment().Init(
		payments.InitRequest{
			Amount:      300,
			CustomerKey: "test_user",
			OrderId:     fmt.Sprint(time.Now().Unix()),
			Receipt: payments.Receipt{
				Taxation: "usn_income",
				Email:    "papkovda@me.com",
				Items: []payments.Item{
					{
						Name:     "Товар 1",
						Price:    100,
						Quantity: 3,
						Amount:   300,
						Tax:      "none",
					},
				},
			},
		})

	if err != nil {
		t.Fatalf("Err: %v", err)
		return
	}

	fmt.Sprintf(resp.PaymentURL)
}

func TestSafeDeal_PaymentInfo(t *testing.T) {
	//3133285872

	sd, err := tinkoff.New(false, lgr.New(lgr.Debug)).SafeDeal(
		"1683019138816",
		"pf1h19bhd7ozster",
	)

	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := sd.Payment().Cancel(payments.CancelRequest{
		PaymentId: "3133285872",
		Amount:    300,
		Receipt: payments.Receipt{
			Taxation: "usn_income",
			Email:    "papkovda@me.com",
			Items: []payments.Item{
				{
					Name:     "Товар 1",
					Price:    100,
					Quantity: 3,
					Amount:   300,
					Tax:      "none",
				},
			},
		},
	})

	if err != nil {
		t.Fatalf("Err: %v", err)
		return
	}

	fmt.Sprintf("%v", resp)
}
