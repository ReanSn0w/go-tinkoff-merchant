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
		t.Fail()
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
		FullName:          "ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ \"ЭЙС ПЛЭЙС\"",
		Name:              "ООО \"ЭЙС ПЛЭЙС\"",
		Inn:               "7743286296",
		Kpp:               "774301001",
		Ogrn:              5187746009810,
		Smz:               false,
		Addresses: []partner.Address{{
			Type:    "legal",
			Zip:     "125212",
			Country: "RUS",
			City:    "Москва",
			Street:  "улица Адмирала Макарова, 15",
		}},
		Email: "llc.aceplace@yandex.ru",
		Ceo: partner.Ceo{
			FirstName:  "Сергей",
			LastName:   "Сухачев",
			MiddleName: "Сергеевич",
			BirthDate:  "1990-03-29",
		},
		SiteURL: "https://aceplace.ru/aceplace",
		BankAccount: partner.BankAccount{
			Account:    "40817810100002965390",
			KorAccount: "30101810145250000974",
			BankName:   "АО «Тинькофф Банк»",
			Bik:        "044525974",
			Details:    "Перевод средств по договору No 0-0 от 16.09.2021 по Реестру Операций от ${date}. Сумма комиссии ${rub} руб. ${kop} коп., НДС не облагается.",
			Tax:        10,
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
