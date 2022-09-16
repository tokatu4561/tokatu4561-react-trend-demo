package main

import (
	"encoding/json"
	"myapp/internal/cards"
	"net/http"
	"strconv"
)

type stripePayload struct {
	Amount        string `json:"amount"`
}

type jsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
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
		Currency: "jpy",
	}

	success := true

	pi, msg, err := card.Charge("jpy", amount)
	if err != nil {
		success = false
	}

	if success {
		out, err := json.MarshalIndent(pi, "", "   ")
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		response := jsonResponse {
			Success: false,
			Message: msg,
			Content: "",
		}

		out, err := json.MarshalIndent(response, "", "   ")
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}
}

func (app *application) CreateAuthnicateToken(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		email string `json:"email"`
		password string `json:"password"`
	}

	err := app.readJson(w, r, &userInput)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// Userを取得
	user, err := app.DB.GetByEmail(userInput.email)
	if err != nil {
		app.badRequest(w, r, err)
	}

	// パスワード検証
	isValidPassword, err := app.passwordMatches(user.Password, userInput.password)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// token生成
}