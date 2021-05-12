package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetAllMonths(w http.ResponseWriter, r *http.Request) {

	// Slice to store the response content
	var res []models.MonthSelector

	// Declare all the expected results variables in order
	var (
		month string
	)

	// Query on database using the stmt SQL and passing two datetime into strings to it
	rows, err := setup.DB.Query(`
		SELECT strftime('%Y-%m-01', d.time_stamp) as Month
		FROM dates d
		GROUP BY strftime('%Y-%m-01', Month)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section(good practice)
	defer rows.Close()

	// Loops over the query results
	for rows.Next() {
		err := rows.Scan(&month)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(id, name)
		// Generate a single Bill struct
		item := models.MonthSelector{
			Month: month,
		}
		// Append this Bill struct to the response slice
		res = append(res, item)
	}
	// Check for any erros on rolls
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Generates an Json based on the response slice
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write output
	w.Write(js)
}
