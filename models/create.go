package models

import "time"

type Companies struct {
	ID   uint   `json:"value"`
	Name string `json:"label"`
}

type Tags struct {
	ID   uint   `json:"value"`
	Name string `json:"label"`
}

type Dates struct {
	ID        uint      `json:"id"`
	TimeStamp time.Time `json:"time_stamp"`
}

type Bills struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	CompanyID uint    `json:"company_id"`
	TagID     uint    `json:"tag_id"`
	DateID    uint    `json:"date_id"`
}

type BillID struct {
	ID uint `json:"id"`
}

type Res struct {
	Title        string `json:"title"`
	Price        string `json:"price"`
	IsNewCompany bool   `json:"isnewCompany"`
	Company      string `json:"company"`
	IsNewTag     bool   `json:"isNewTag"`
	Tag          string `json:"tag"`
	Date         string `json:"date"`
}
