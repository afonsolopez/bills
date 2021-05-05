package models

import "time"

type Bill struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Company   string    `json:"company"`
	Tag       string    `json:"tag"`
	TimeStamp time.Time `json:"timeStamp"`
}
