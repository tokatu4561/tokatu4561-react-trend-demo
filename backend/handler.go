package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/internal/cards"
	"myapp/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

type stripePayload struct {
	Amount string `json:"amount"`
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
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
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
		response := jsonResponse{
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

// TODO: csrf対策する(csrfトークン)
func (app *application) CreateAuthnicateToken(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &userInput)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}

	// Userを取得
	user, err := app.Models.User.GetByEmail(userInput.Email)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}

	// パスワード検証
	validPassword, err := app.passwordMatches(user.Password, userInput.Password)
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w, r, err)
		return
	}

	if !validPassword {
		app.badRequest(w, r, errors.New("パスワードが一致しません"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// TODO:: エラーログ残す　後でやる
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		HttpOnly: true,
	}

	var resp struct {
		Error bool `json:"error"`
	}
	resp.Error = false

	http.SetCookie(w, cookie)
	app.writeJson(w, 200, resp)
}

func (app *application) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")

	// payloadToken, err := app.getAuthenticateTokenFromHeader(r)
	payloadToken := cookie.Value
	if err != nil {
		app.badRequest(w, r, err)
	}

	//　tokenを検証
	token, _ := jwt.Parse(payloadToken, func(token *jwt.Token) (interface{}, error) {
		// tokenアルゴリズムを検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// TODO:WARNING: シークレットキーはenvで管理したい
		return []byte("SECRET_KEY"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var response struct {
			Error bool        `json:"error"`
			User  models.User `json:"user_id"`
		}

		userId := claims["user_id"].(float64)
		user, err := app.Models.User.GetById(int(userId))
		if err != nil {
			app.errorLog.Fatalln(err)
			app.badRequest(w, r, err)
			return
		}

		response.Error = false
		response.User = *user

		app.writeJson(w, 200, response)
	} else {
		app.badRequest(w, r, errors.New("認証エラー"))
	}
}

// GetSale returns existing all Task as json
func (app *application) GetTasks(w http.ResponseWriter, r *http.Request) {
}

// GetSale returns one Todo as json
func (app *application) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taskID, _ := strconv.Atoi(id)

	task, err := app.Models.Task.GetById(taskID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, task)
}

// CreateTask insert task ind db and returns Task as json
func (app *application) CreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	userID, _ := strconv.Atoi(r.Form.Get("user_id"))
	task := models.Task{
		Title: r.Form.Get("title"),
	}

	newTask, err := app.Models.Task.Insert(userID, task)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	app.writeJson(w, http.StatusCreated, newTask)
}

func (app *application) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := app.readJson(w, r, task)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	err = app.Models.Task.Update(task)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, nil)
}

func (app *application) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskID int
	err := app.readJson(w, r, taskID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	err = app.Models.Task.Delete(taskID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	app.writeJson(w, http.StatusOK, nil)
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
