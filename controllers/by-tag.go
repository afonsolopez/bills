package controllers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/afonsolopez/bills/functions"
	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetMonthBillsByTag(w http.ResponseWriter, r *http.Request) {

	var res []models.ByTag

	_, firstDay, lastDay := functions.ConditionalDate(w, r)

	// Declare all the expected results variables in order
	var (
		tag       string
		total     float64
		timeStamp time.Time
	)

	// Make a query on database
	stmt, err := setup.DB.Prepare(`
		SELECT t.name AS tag, SUM(b2.price) AS total, d2.time_stamp 
		FROM bills b2 
		LEFT JOIN tags t
		ON b2.tag_id = t.id
		LEFT JOIN dates d2 
		ON b2.date_id = d2.id 
		WHERE d2.time_stamp
		BETWEEN ? AND ?
		GROUP BY t.name;
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section (good practice)
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
		err := rows.Scan(&tag, &total, &timeStamp)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(tag, total)
		// Generate a single Bill struct
		item := models.ByTag{
			Tag:   tag,
			Total: math.Ceil(total*100) / 100,
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
