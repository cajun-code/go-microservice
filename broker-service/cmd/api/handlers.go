package main

import (
	"net/http"
	"web-helpers/helpers"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	helpers.WriteJSON(w, http.StatusOK, payload)
}
