package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func AfterCreate(str string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// CreateNewBill ...
func CreateNewBill(w http.ResponseWriter, r *http.Request) {

	var (
		// allDates  []models.Dates
		// title     string
		// price     float64

		companyID   uint
		companyName string

		tagID   uint
		tagName string

		dateID        uint
		dateTimeStamp string

		cID   uint
		cName string

		tID   uint
		tName string

		dID        uint
		dTimeStamp string
	)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg models.Res
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	convertedPrice, err := strconv.ParseFloat(msg.Price, 64)
	if err != nil {
		fmt.Println("Coul not convert string to float64")
		http.Error(w, err.Error(), 500)
		return
	}

	//=============================================================================
	// New company handler
	//=============================================================================

	if msg.IsNewCompany {
		// CREATE NEW COMPANY AND USE IT ID

		fmt.Println("This insert will create a new company record")

		sqlStatement := `
			SELECT c.id, c.name 
			FROM companies c
			WHERE c.name=?
	`

		// Query for a company with same name as received
		row := setup.DB.QueryRow(sqlStatement, msg.Company)
		err := row.Scan(&cID, &cName)

		if err != nil {
			// If there's no match for that company, insert a new one
			if err == sql.ErrNoRows {
				fmt.Println("Zero company rows found, inserting new company name...")
				// Execute SQL insert on companies table
				_, err = setup.DB.Exec(`
				INSERT INTO companies (name)
				VALUES (?)
				`, msg.Company)

				if err != nil {
					panic(err)
				}

			} else {
				panic(err)
			}
		}

		// After the insertion, query that result to get it ID
		queryNewCompany := `
			SELECT c.id, c.name 
			FROM companies c
			WHERE c.name=?
		`
		row2 := setup.DB.QueryRow(queryNewCompany, msg.Company)
		err2 := row2.Scan(&cID, &cName)

		if err2 != nil {
			if err2 == sql.ErrNoRows {
				fmt.Println("Zero company rows found")
			} else {
				panic(err)
			}
		}

		// Attribute the new inserted comany name to the variable
		companyID = cID
		companyName = cName
		fmt.Println("New value added to company of: ", companyName)

	} else {
		// ### AD SEARCH BY COMPANY NAME AND RETURN COMPANY ID ###
		fmt.Println("This insert will use an older company record")

		queryNewCompany := `
		SELECT c.id, c.name 
		FROM companies c
		WHERE c.name=?
	`
		row2 := setup.DB.QueryRow(queryNewCompany, msg.Company)
		err2 := row2.Scan(&cID, &cName)

		if err2 != nil {
			if err2 == sql.ErrNoRows {
				fmt.Println("Zero company rows found")
			} else {
				panic(err)
			}
		}

		// Attribute the new inserted comany name to the variable
		companyID = cID
		companyName = cName
		fmt.Println("New value added to company of: ", companyName)
	}

	//=============================================================================
	// New tag handler
	//=============================================================================

	if msg.IsNewTag {
		fmt.Println("This insert will create a new Tag record")

		sqlStatement := `
			SELECT t.id, t.name 
			FROM tags t
			WHERE t.name=?
	`

		// Query for a company with same name as received
		row := setup.DB.QueryRow(sqlStatement, msg.Tag)
		err := row.Scan(&tID, &tName)

		if err != nil {
			// If there's no match for that company, insert a new one
			if err == sql.ErrNoRows {
				fmt.Println("Zero tag rows found, inserting new company name...")

				// Execute an SQL insert on tags table
				_, err = setup.DB.Exec(`
				INSERT INTO tags (name)
				VALUES (?)
				`, msg.Tag)

				if err != nil {
					panic(err)
				}

			} else {
				panic(err)
			}
		}

		// After the insertion, query that result to get it ID
		queryNewTag := `
			SELECT t.id, t.name 
			FROM tags t
			WHERE t.name=?
		`
		row2 := setup.DB.QueryRow(queryNewTag, msg.Tag)
		err2 := row2.Scan(&tID, &tName)

		if err2 != nil {
			if err2 == sql.ErrNoRows {
				fmt.Println("Zero tag rows found")
			} else {
				panic(err)
			}
		}

		tagID = tID
		tagName = tName
		fmt.Println("New value added to tag of: ", tagName)

	} else {
		fmt.Println("This insert will use an older Tag record")
		queryNewTag := `
		SELECT t.id, t.name 
		FROM tags t
		WHERE t.name=?
	`
		row2 := setup.DB.QueryRow(queryNewTag, msg.Tag)
		err2 := row2.Scan(&tID, &tName)

		if err2 != nil {
			if err2 == sql.ErrNoRows {
				fmt.Println("Zero tag rows found")
			} else {
				panic(err)
			}
		}

		tagID = tID
		tagName = tName
		fmt.Println("New value added to tag of: ", tagName)

	}

	//=============================================================================
	// New date handler
	//=============================================================================

	queryDate := `
	SELECT id, time_stamp 
	FROM dates d2
	WHERE time_stamp = ?
`
	rowD := setup.DB.QueryRow(queryDate, msg.Date)
	errD := rowD.Scan(&dID, &dTimeStamp)

	if errD != nil {
		if errD == sql.ErrNoRows {
			fmt.Println("Zero date rows found, inserting new time stamp...")

			// Execute an SQL insert on tags table
			_, err = setup.DB.Exec(`
			INSERT INTO dates (time_stamp)
			VALUES (?)
			`, msg.Date)

			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// After the insertion, query that result to get it ID
	queryNewDate := `
			SELECT d.id, d.time_stamp 
			FROM dates d
			WHERE d.time_stamp=?
		`
	rowD2 := setup.DB.QueryRow(queryNewDate, msg.Date)
	errD2 := rowD2.Scan(&dID, &dTimeStamp)

	if errD2 != nil {
		if errD2 == sql.ErrNoRows {
			fmt.Println("Zero tag rows found")
		} else {
			panic(err)
		}
	}

	dateID = dID
	dateTimeStamp = dTimeStamp
	fmt.Println("New value added to date of: ", dateTimeStamp)

	//=============================================================================
	// New bill handler
	//=============================================================================

	insert := &models.Bills{
		Title:     msg.Title,
		Price:     convertedPrice,
		CompanyID: companyID,
		TagID:     tagID,
		DateID:    dateID,
	}

	// Execute SQL insert on bills table
	_, err = setup.DB.Exec(`
	INSERT INTO bills
	(title, price, company_id, tag_id, date_id)
	VALUES(?, ?, ?, ?, ?);
	`, msg.Title, convertedPrice, companyID, tagID, dateID)

	if err != nil {
		panic(err)
	}

	// Generates an Json based on the response slice
	js, err := json.Marshal(insert)
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
