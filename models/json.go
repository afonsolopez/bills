package models

import (
	"encoding/json"
	"math"
	"time"

	"github.com/afonsolopez/bills/functions"
)

type JsonBill struct {
	Company string  `json:"company"`
	Price   float64 `json:"price"`
	Tag     string  `json:"tag"`
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
	return math.Ceil(sum*100) / 100
}

func (r *Response) RemainingDays() int {
	// Gets the actual timestamp
	now := time.Now()
	monthWorker := functions.MonthWorker{Month: functions.Month{Day: now, Gap: 1}}
	// Gets the difference between two dates
	diff := monthWorker.LastOfMonth().Sub(now)
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
