package functions

import (
	"fmt"
	"time"
)

func ParseDate(str string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
