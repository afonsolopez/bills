package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

var bills []models.Bill

// GetLastMonthBills ...
func GetLastMonthBills(w http.ResponseWriter, r *http.Request) {

	// // Gets the actual timestamp
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	// Gets the actual location
	currentLocation := now.Location()
	// Finds the first day of the actual month
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	// Query between the range of entire months choosed
	setup.DB.Joins("Company").Joins("Tag").Joins("Date").Where("Date__time_stamp BETWEEN ? AND ?", firstOfMonth, now).Order("Date__time_stamp desc").Find(&bills)
	// Variable to store slice of processed bills
	var getBills []models.JsonBill
	// Process each bill and inject it to "getBills"
	for _, b := range bills {
		r := models.JsonBill{
			Company: b.Company.Name,
			Price:   b.Price,
			Year:    b.Date.TimeStamp.Year(),
			Month:   int(b.Date.TimeStamp.Month()),
			Day:     b.Date.TimeStamp.Day(),
		}
		getBills = append(getBills, r)
	}

	// Generates a response struct and inject all the processed bill on it
	res := models.Response{
		Bills: getBills,
	}

	// Return JSON encoding to output
	output, err := json.Marshal(res)

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
