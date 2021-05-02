package models

import "time"

type ByTag struct {
	Name  string    `json:"name"`
	Total float64   `json:"price"`
	Time  time.Time `json:"timeStamp"`
}
