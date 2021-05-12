package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {

	// Slice to store the response content
	var res []models.Companies

	// Declare all the expected results variables in order
	var (
		id   uint
		name string
	)

	// Query on database using the stmt SQL and passing two datetime into strings to it
	rows, err := setup.DB.Query(`
		SELECT DISTINCT id, name
		FROM companies c2 
		ORDER BY name ASC
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Closes the rows variable section(good practice)
	defer rows.Close()

	// Loops over the query results
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(id, name)
		// Generate a single Bill struct
		item := models.Companies{
			ID:   id,
			Name: name,
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
