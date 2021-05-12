package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"
)

func DeleteBill(w http.ResponseWriter, r *http.Request) {

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg models.BillID
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//=============================================================================
	// Query bill info
	//=============================================================================

	// Declare all the expected results variables in order
	var (
		id        uint
		title     string
		price     float64
		companyID uint
		tagID     uint
		dateID    uint
	)

	queryBill := `
	SELECT b2.id, b2.title, b2.price, b2.company_id, b2.tag_id, b2.date_id 
	FROM bills b2 
	WHERE b2.id = ?
		`
	rowB := setup.DB.QueryRow(queryBill, msg.ID)
	errB := rowB.Scan(&id, &title, &price, &companyID, &tagID, &dateID)

	if errB != nil {
		if errB == sql.ErrNoRows {
			fmt.Println("Zero company rows found")
		} else {
			panic(err)
		}
	}

	//=============================================================================
	// Delete bill
	//=============================================================================

	delRes := struct {
		Msg    string `json:"msg"`
		BillID uint   `json:"bill_id"`
	}{
		Msg:    "Bill, deleted with success.",
		BillID: msg.ID,
	}

	sqlStatement := `
	DELETE FROM bills
	WHERE id = ?;`
	_, err = setup.DB.Exec(sqlStatement, msg.ID)
	if err != nil {
		panic(err)
	}

	//=============================================================================
	// Query tag info
	//=============================================================================

	// Declare all the expected results variables in order
	var (
		tags int
	)

	queryTag := `
	SELECT COUNT(b2.tag_id) AS tags
	FROM bills b2 
	WHERE b2.tag_id = ?
		`
	rowT := setup.DB.QueryRow(queryTag, tagID)
	errT := rowT.Scan(&tags)

	if errT != nil {
		if errT == sql.ErrNoRows {
			fmt.Println("Zero tag rows found")
		} else {
			panic(err)
		}
	}

	//=============================================================================
	// Query company info
	//=============================================================================

	// Declare all the expected results variables in order
	var (
		companies int
	)

	queryCompany := `
	SELECT COUNT(b2.company_id) AS companies
	FROM bills b2 
	WHERE b2.company_id = ?
		`
	rowC := setup.DB.QueryRow(queryCompany, companyID)
	errC := rowC.Scan(&companies)

	if errC != nil {
		if errC == sql.ErrNoRows {
			fmt.Println("Zero tag rows found")
		} else {
			panic(err)
		}
	}

	//=============================================================================
	// Query date info
	//=============================================================================

	// Declare all the expected results variables in order
	var (
		dates int
	)

	queryDate := `
	SELECT COUNT(b2.date_id) AS dates
	FROM bills b2 
	WHERE b2.date_id = ?
		`
	rowD := setup.DB.QueryRow(queryDate, dateID)
	errD := rowD.Scan(&dates)

	if errD != nil {
		if errD == sql.ErrNoRows {
			fmt.Println("Zero date rows found")
		} else {
			panic(err)
		}
	}

	//=============================================================================
	// Delete tag/company/date if needed
	//=============================================================================

	if tags == 0 {
		deleteTag := `
			DELETE FROM tags
			WHERE id = ?;`
		_, err = setup.DB.Exec(deleteTag, tagID)
		if err != nil {
			panic(err)
		}
	}

	if companies == 0 {
		deleteCompany := `
			DELETE FROM companies
			WHERE id = ?;`
		_, err = setup.DB.Exec(deleteCompany, companyID)
		if err != nil {
			panic(err)
		}
	}

	if dates == 0 {
		deleteDate := `
			DELETE FROM dates
			WHERE id = ?;`
		_, err = setup.DB.Exec(deleteDate, dateID)
		if err != nil {
			panic(err)
		}
	}

	// Generates an Json based on the response slice
	js, err := json.Marshal(delRes)
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
