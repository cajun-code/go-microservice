package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web-helpers/helpers"

	"github.com/pkg/errors"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	helpers.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.ErrorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
		return
	default:
		helpers.ErrorJSON(w, errors.New("unknown action"))
		return
	}
	//app.authenticate(w, requestPayload.Auth)
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we'll send in the response
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	request, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		helpers.ErrorJSON(w, err)
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		helpers.ErrorJSON(w, err)
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		fmt.Println("invalid credentials")
		return
	} else if response.StatusCode != http.StatusAccepted {
		helpers.ErrorJSON(w, errors.New("Error calling auth service"), http.StatusUnauthorized)
		fmt.Println("Error calling auth service")
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService helpers.JSONResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		helpers.ErrorJSON(w, err)
		fmt.Println(err)
		return
	}
	if jsonFromService.Error {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	fmt.Println("Authenticated!")
	var payload = helpers.JSONResponse{
		Error:   false,
		Message: "Authenticated!",
		Data:    jsonFromService.Data,
	}
	helpers.WriteJSON(w, http.StatusAccepted, payload)
}
