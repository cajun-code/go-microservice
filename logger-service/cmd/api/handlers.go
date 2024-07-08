package main

import (
	"log-service/data"
	"net/http"
	"web-helpers/helpers"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var payload JSONPayload
	_ = helpers.ReadJSON(w, r, &payload)

	event := data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}

	err := app.Models.LogEntry.Insert(event)

	if err != nil {
		helpers.ErrorJSON(w, err)
		return
	}
	response := helpers.JSONResponse{
		Error:   false,
		Message: "logged",
	}

	helpers.WriteJSON(w, http.StatusAccepted, response)
}
