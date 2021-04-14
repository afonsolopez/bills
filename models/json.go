package models

import (
	"encoding/json"
	"time"
)

type JsonBill struct {
	Company string  `json:"company"`
	Price   float64 `json:"price"`
	Year    int     `json:"year"`
	Month   int     `json:"month"`
	Day     int     `json:"day"`
}

type Response struct {
	Bills []JsonBill `json:"bills"`
}

func (r *Response) TotalValue() float64 {
	sum := 0.00
	for _, i := range r.Bills {
		sum += i.Price
	}
	return sum
}

func (r *Response) RemainingDays() int {
	// Gets the actual timestamp
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	// Gets the actual location
	currentLocation := now.Location()
	// Finds the first day of the actual month
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	// Finds the last day of the month
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	// Gets the difference between two dates
	diff := lastOfMonth.Sub(now)
	return int(diff.Hours() / 24) // number of days

}

// Custom marshal with a temporary struct
func (r Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Bills         []JsonBill `json:"bills"`
		Total         float64    `json:"total"`
		RemainingDays int        `json:"remainingDays"`
	}{
		Bills:         r.Bills,
		Total:         r.TotalValue(),
		RemainingDays: r.RemainingDays(),
	})
}
