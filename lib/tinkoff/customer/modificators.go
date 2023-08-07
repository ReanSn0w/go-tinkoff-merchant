package customer

type RequestModificator func(Request) Request

func AddPhone(phone string) RequestModificator {
	return func(r Request) Request {
		r.Phone = phone
		return r
	}
}

func AddEmail(email string) RequestModificator {
	return func(r Request) Request {
		r.Email = email
		return r
	}
}

func WithCardID(cardID string) RequestModificator {
	return func(r Request) Request {
		r.CardId = cardID
		return r
	}
}
