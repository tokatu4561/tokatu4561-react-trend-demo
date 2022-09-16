package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func(app *application) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func(app *application) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) error {
	var payload struct {
		Error bool `json:"error"`
		Message string `json:"message"`
	}

	payload.Error = true
	payload.Message = err.Error()

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)

	return nil
}

// hash化されている password と plainTextが一致しているかを確認する
func (app *application) passwordMatches(password string, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}