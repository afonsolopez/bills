package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/afonsolopez/bills/functions"
	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetThisMonthBills(w http.ResponseWriter, r *http.Request) {

	var res []models.Bill

	// Declare all the expected results variables in order
	var (
		id         int
		title      string
		price      float64
		company    string
		tag        string
		time_stamp time.Time
	)

	// Get today's date
	now := time.Now()
	// Get the first day of the current month date
	monthWorker := functions.MonthWorker{Month: functions.Month{Day: now, Gap: 1}}
	firstOfMonth := monthWorker.FirstOfMonth()

	// Make a query on database
	stmt, err := setup.DB.Prepare(`
		SELECT b2.id, b2.title, b2.price, c2.name AS company, t.name AS tag, d2.time_stamp 
		FROM bills b2 
		LEFT JOIN companies c2 
		ON b2.company_id = c2.id
		LEFT JOIN tags t
		ON b2.tag_id = t.id 
		LEFT JOIN dates d2 
		ON b2.date_id = d2.id 
		WHERE d2.time_stamp
		BETWEEN ? AND ?
		ORDER BY d2.time_stamp DESC
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section (good practice)
	defer stmt.Close()

	// Query on database using the stmt SQL and passing two datetime into strings to it
	rows, err := stmt.Query(firstOfMonth.String(), now.String())
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section(good practice)
	defer rows.Close()

	// Loops over the query results
	for rows.Next() {
		err := rows.Scan(&id, &title, &price, &company, &tag, &time_stamp)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title, price, company, tag, time_stamp)
		// Generate a single Bill struct
		item := models.Bill{
			Id:        id,
			Title:     title,
			Price:     price,
			Company:   company,
			Tag:       tag,
			TimeStamp: time_stamp,
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
