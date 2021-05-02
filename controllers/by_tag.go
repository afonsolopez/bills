package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afonsolopez/bills/functions"
	"github.com/afonsolopez/bills/setup"
)

// var by []models.ByTag

// GetLastMonthBills ...
func GetMonthBillsByTag(w http.ResponseWriter, r *http.Request) {
	// Get today's date
	now := time.Now()
	// Get the first day of the current month date
	monthWorker := functions.MonthWorker{Month: functions.Month{Day: now, Gap: 1}}
	firstOfMonth := monthWorker.FirstOfMonth()
	// Query between the range of entire months choosed
	setup.DB.Joins("Company").Joins("Tag").Joins("Date").Joins("Bills").Select("tag.name, date.time_stamp as time, sum(bills.price) as total").Where("date.time_stamp BETWEEN ? AND ?", firstOfMonth, now).Group("tag.name").Order("date.time_stamp desc").Find(&bills)
	// Variable to store slice of processed bills
	// var getBills []models.ByTag
	// Process each bill and inject it to "getBills"
	// for _, b := range bills {
	// 	r := models.ByTag{
	// 		Tag:   "",
	// 		Price: 0,
	// 		Year:  0,
	// 		Month: 0,
	// 	}
	// 	getBills = append(getBills, r)
	// }

	// Generates a response struct and inject all the processed bill on it
	// res := models.Response{
	// 	Bills: getBills,
	// }

	// Return JSON encoding to output
	output, err := json.Marshal(&bills)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write output
	w.Write([]byte(output))

}
