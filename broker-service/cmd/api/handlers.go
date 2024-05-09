package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	app.writeJSON(w, http.StatusOK, payload)
}