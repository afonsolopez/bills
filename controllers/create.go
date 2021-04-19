package controllers

import (
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

	// setup.DB.Where("name = ?", msg.Company).First(&companies)

	price, err := strconv.ParseFloat(msg.Price, 64)
	if err != nil {
		fmt.Println("Coul not convert string to float64")
		http.Error(w, err.Error(), 500)
		return
	}

	insert := &models.Bill{
		Title:  msg.Title,
		Price:  price,
		DateID: 0,
		Date:   models.Date{TimeStamp: AfterCreate(msg.Date)},
	}

	if msg.IsNewCompany {
		fmt.Println("This insert will create a new company record")
		insert.Company = models.Company{Name: msg.Company}
	} else {
		fmt.Println("This insert will use an older company record")

		setup.DB.First(&companies, "name = ?", msg.Company)

		var getCompanyId []models.Companies

		for _, c := range companies {
			r := models.Companies{
				ID:   c.ID,
				Name: c.Name,
			}
			getCompanyId = append(getCompanyId, r)
		}
		fmt.Println(getCompanyId)

		insert.CompanyID = int(getCompanyId[0].ID)
	}

	if msg.IsNewTag {
		fmt.Println("This insert will create a new Tag record")
		insert.Tag = models.Tag{Name: msg.Tag}
	} else {
		fmt.Println("This insert will use an older Tag record")

		setup.DB.First(&tags, "name = ?", msg.Tag)

		var getTagId []models.Tags

		for _, c := range tags {
			r := models.Tags{
				ID:   c.ID,
				Name: c.Name,
			}
			getTagId = append(getTagId, r)
		}
		fmt.Println(getTagId)

		insert.TagID = int(getTagId[0].ID)
	}

	setup.DB.Create(&insert)

	// Marshal
	output, err := json.Marshal(&insert)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header Content-Type
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Credentials", "true")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	// Write output
	w.Write(output)

}
