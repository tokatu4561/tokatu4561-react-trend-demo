package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/internal/cards"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
		return
	}

	// パスワード検証
	isValidPassword, err := app.passwordMatches(user.Password, userInput.password)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	if (isValidPassword) {
		app.badRequest(w, r, errors.New("パスワードが一致しません"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": 12345678,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			})
	
	// TODO:: エラーログ残す　後でやる
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))

	var response struct {
		Error bool `json:"error"`
		Token string `json:"token"`
	}

	response.Error = false
	response.Token = tokenString

	app.writeJson(w, 200, response)
}

func (app *application) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	payloadToken, err := app.getAuthenticateTokenFromHeader(r)
	if err != nil {
		app.badRequest(w, r, err)
	}

	//　tokenを検証
	token, _ := jwt.Parse(payloadToken, func(token *jwt.Token) (interface{}, error) {
		// tokenアルゴリズムを検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRET_KEY"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims); 
	if ok && token.Valid {
		var response struct {
			Error bool `json:"error"`
			UserId float64 `json:"user_id"`
		}
	
		response.Error = false
		response.UserId = claims["user_id"].(float64)
	
		app.writeJson(w, 200, response)
	} else {
		app.badRequest(w, r, errors.New("認証エラー"))
	}
}

// ヘッダーのbearerからトークンを取り出す
func (app *application) getAuthenticateTokenFromHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("no authorization header received")
	}

	token := headerParts[1]
	if len(token) != 26 {
		return "", errors.New("authentication token wrong size")
	}

	return token, nil
}