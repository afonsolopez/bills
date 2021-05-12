package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/afonsolopez/bills/models"
)

func ConditionalDate(w http.ResponseWriter, r *http.Request) (MonthWorker, string, string) {
	var (
		monthWorker MonthWorker
		firstDay    string
		lastDay     string
		day         time.Time
	)

	if r.Method == http.MethodGet {
		// If GET request it gets the first day of the month and actual day

		// Get today's date
		day = time.Now()

		// Get the first day of the current month date
		monthWorker = MonthWorker{Month: Month{Day: day, Gap: 1}}
		firstOfMonth := monthWorker.FirstOfMonth()

		firstDay = firstOfMonth.String()[0:10]
		lastDay = day.String()[0:10]

		return monthWorker, firstDay, lastDay

	} else if r.Method == http.MethodPost {
		// If POST request get first and last day of that month

		// Read body
		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("body bug")
			http.Error(w, err.Error(), 500)
			return monthWorker, "", ""
		}

		// Unmarshal
		var msg models.DateQuery
		err = json.Unmarshal(b, &msg)
		if err != nil {
			fmt.Println("unmarshal bug")
			http.Error(w, err.Error(), 500)
			return monthWorker, "", ""
		}
		fmt.Println(msg.Year, msg.Month)
		stringfyDay := msg.Year + "-" + msg.Month + "-01"

		day = ParseDate(stringfyDay)

		// Get the first day of the current month date
		monthWorker = MonthWorker{Month: Month{Day: day, Gap: 1}}
		firstOfMonth := monthWorker.FirstOfMonth()
		lastOfMonth := monthWorker.LastOfMonth()

		firstDay = firstOfMonth.String()[0:10]
		lastDay = lastOfMonth.String()[0:10]

		return monthWorker, firstDay, lastDay
	}
	return monthWorker, "", ""
}
