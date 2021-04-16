package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

var tags []models.Tag

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	setup.DB.Distinct("name").Find(&tags)

	var responses []models.AllTags

	for _, t := range tags {
		r := models.AllTags{
			Label: t.Name,
			Value: t.Name,
		}
		responses = append(responses, r)
	}
	fmt.Println(responses)
	// Return JSON encoding to output
	output, err := json.Marshal(responses)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write output
	w.Write(output)
}
