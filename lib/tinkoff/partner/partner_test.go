package partner_test

import (
	"testing"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff"
	"github.com/ReanSn0w/go-tinkoff-merchant/lib/tinkoff/partner"
	"github.com/go-pkgz/lgr"
)

var (
	service = tinkoff.New(true, lgr.Default())
)

func Test_New(t *testing.T) {
	_, err := service.Partner("aceplace", "aceplace")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestPartner_Register(t *testing.T) {
	p, err := service.Partner("aceplace", "aceplace")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	ID := "someuserid"

	resp, err := p.Register(partner.RegistrationRequest{
		ID:                ID,
		BillingDescriptor: "AcePlace",
		FullName:          "ООО ЭйсПлейс",
		Name:              "ООО ЭйсПлейс",
		Inn:               "",
		Kpp:               "",
		Ogrn:              0,
		Smz:               false,
		Addresses: []partner.Address{{
			Type:    "legal",
			Zip:     "",
			Country: "RU",
			City:    "Москва",
			Street:  "улица Адмирала Макарова, 15",
		}},
		Email: "llc.aceplace@yandex.ru",
		Ceo: partner.Ceo{
			FirstName:  "",
			LastName:   "",
			MiddleName: "",
			BirthDate:  "",
			//Country:    "",
		},
		SiteURL: "https://aceplace.ru/aceplace",
		BankAccount: partner.BankAccount{
			Account:    "",
			KorAccount: "",
			BankName:   "",
			Bik:        "",
			Kbk:        "",
			Oktmo:      "",
			Details:    "",
			Tax:        0,
		},
	})

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if resp.ID != ID || resp.ShopCode == 0 {
		t.Log("Не удалось получить данные партнера")
		t.FailNow()
	}
}

func TestPartner_Update(t *testing.T) {
	p, err := service.Partner("aceplace", "aceplace")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	ID := "someuserid"

	resp, err := p.Update(ID, partner.BankAccount{
		Account:    "",
		KorAccount: "",
		BankName:   "",
		Bik:        "",
		Kbk:        "",
		Oktmo:      "",
		Details:    "",
		Tax:        0,
	})

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if resp.ID != ID || resp.ShopCode == 0 {
		t.Log("Не удалось получить данные партнера")
		t.FailNow()
	}
}
