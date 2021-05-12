package controllers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/afonsolopez/bills/functions"
	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetMonthBillsByDay(w http.ResponseWriter, r *http.Request) {

	// Slice to store the response content
	var res []models.ByDay

	// Declare all the expected results variables in order
	var (
		timeStamp time.Time
		total     float64
	)

	// Get today's date
	now := time.Now()
	// Get the first day of the current month date
	monthWorker := functions.MonthWorker{Month: functions.Month{Day: now, Gap: 1}}
	firstOfMonth := monthWorker.FirstOfMonth()

	// Prepare a query on database with two datetime arguments
	stmt, err := setup.DB.Prepare(`
		SELECT d2.time_stamp AS time_stamp, SUM(b2.price) AS total
		FROM bills b2 
		LEFT JOIN dates d2 
		ON b2.date_id = d2.id
		WHERE d2.time_stamp
		BETWEEN ? AND ?
		GROUP BY time_stamp;
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the stmt variable section (good practice)
	defer stmt.Close()

	// Query on database using the stmt SQL and passing two datetime into strings to it
	rows, err := stmt.Query(firstOfMonth.String()[0:10], now.String()[0:10])
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section(good practice)
	defer rows.Close()

	// Loops over the query results
	for rows.Next() {
		err := rows.Scan(&timeStamp, &total)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(timeStamp, total)
		// Generate a single Bill struct
		item := models.ByDay{
			TimeStamp: strconv.Itoa(timeStamp.Day()),
			Total:     math.Ceil(total*100) / 100,
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
