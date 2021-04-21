package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

// var bills []models.Bill

func GetAllBills(w http.ResponseWriter, r *http.Request) {
	setup.DB.Joins("Company").Joins("Tag").Joins("Date").Order("Date__time_stamp desc").Limit(10).Find(&bills)

	// Create Message JSON data
	// message := &bills

	var responses []models.JsonBill

	for _, b := range bills {
		r := models.JsonBill{
			Company: b.Company.Name,
			Price:   b.Price,
			Tag:     b.Tag.Name,
			Year:    b.Date.TimeStamp.Year(),
			Month:   int(b.Date.TimeStamp.Month()),
			Day:     b.Date.TimeStamp.Day(),
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
