package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/afonsolopez/bills/functions"
	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetBigNumbers(w http.ResponseWriter, r *http.Request) {

	// Slice to store the response content
	var res []models.BigNumbers

	// Function to conditionally get the first and last day of a month and they month worker
	monthWorker, firstDay, lastDay := functions.ConditionalDate(w, r)

	// Declare all the expected results variables in order
	var (
		total float64
	)

	// Prepare a query on database with two datetime arguments
	stmt, err := setup.DB.Prepare(`
		SELECT SUM(b2.price) AS total
		FROM bills b2 
		LEFT JOIN dates d2 
		ON b2.date_id = d2.id
		WHERE d2.time_stamp
		BETWEEN ? AND ?
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the stmt variable section (good practice)
	defer stmt.Close()

	// Query on database using the stmt SQL and passing two datetime into strings to it
	rows, err := stmt.Query(firstDay, lastDay)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section(good practice)
	defer rows.Close()

	// Loops over the query results
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			fmt.Println("Total scan equal to:", total) // TEMP FIX - PASS ERR
			// log.Fatal(err)
		}

		// TEMP FIX - NEEDS IMPROVEMENT
		if total == 0 {
			total = 0.0
		}

		// log.Println(remainingDays, total)

		// Generate a single Bill struct
		item := models.BigNumbers{
			RemainingDays: monthWorker.RemainingDays(),
			Total:         total,
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
