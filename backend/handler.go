package main

import (
	"encoding/json"
	"myapp/internal/cards"
	"net/http"
	"strconv"
)

type stripePayload struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}


	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key: app.config.stripe.key,
		Currency: payload.Currency,
	}

	pi, _, err := card.Charge(payload.Currency, amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	out, err := json.MarshalIndent(pi, "", "   ")
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}