package controllers

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/afonsolopez/bills/models"
)

func ShowMessage(w http.ResponseWriter, r *http.Request) {
	// Create Message JSON data
	message := models.Message{
		Text: runtime.Version(),
	}

	// Return JSON encoding to output
	output, err := json.Marshal(message)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write output
	w.Write(output)
}
