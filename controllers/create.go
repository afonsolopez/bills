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

	// if msg.IsNewCompany {
	// 	setup.DB.Create(&models.Company{
	// 		Name: msg.Company,
	// 	})
	// }

	// setup.DB.Where("name = ?", msg.Company).First(&companies)

	// var responses []models.Comp

	// for _, c := range companies {
	// 	r := models.Comp{
	// 		ID:   c.ID,
	// 		Name: c.Name,
	// 	}
	// 	responses = append(responses, r)
	// }
	// fmt.Println(responses)

	price, err := strconv.ParseFloat(msg.Price, 64)
	if err != nil {
		fmt.Println("Coul not convert string to float64")
		http.Error(w, err.Error(), 500)
		return
	}

	setup.DB.Create(&models.Bill{
		// Model:     gorm.Model{},
		CompanyID: 0,
		Company:   models.Company{Name: msg.Company},
		Title:     msg.Title,
		Price:     price,
		DateID:    0,
		Date:      models.Date{TimeStamp: AfterCreate(msg.Date)},
		TagID:     0,
		Tag:       models.Tag{Name: msg.Tag},
	})

	// Marshal
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// json.NewEncoder(w).Encode(companies)

	// Set header Content-Type
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Credentials", "true")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	// Write output
	w.Write(output)

}
