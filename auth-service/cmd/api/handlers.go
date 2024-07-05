package main

import (
	"errors"
	"net/http"
	"web-helpers/helpers"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// validate user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	payload := helpers.JSONResponse{
		Error:   false,
		Message: "logged in",
		Data:    user,
	}
	helpers.WriteJSON(w, http.StatusAccepted, payload)
}
