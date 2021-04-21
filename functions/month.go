package functions

import (
	"math"
	"time"
)

// Month sctruct
type Month struct {
	Day time.Time // Day to work with
	Gap int       // Gap in time by months 1/2/3/12
}

// Create the current Date and Location for that day
func (m *Month) Current() (int, time.Month, *time.Location) {

	// Define current year and month
	currentYear, currentMonth, _ := m.Day.Date()

	// Gets the actual location
	currentLocation := m.Day.Location()

	return currentYear, currentMonth, currentLocation
}

// Struct to handle the original month struct with new methods
type MonthWorker struct {
	Month Month
}

// Get the first day of the actual month
func (m *MonthWorker) FirstOfMonth() time.Time {

	// Define current year, month and location
	year, month, location := m.Month.Current()

	return time.Date(year, month, 1, 0, 0, 0, 0, location)

}

// Get the last day of the actual month
func (m *MonthWorker) LastOfMonth() time.Time {

	// Define current year, month and location
	year, month, location := m.Month.Current()

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

	return firstOfMonth.AddDate(0, 1, -1)

}

// Get the first day of a previous or next month
func (m *MonthWorker) FirstOfLastOfMonth() time.Time {

	// Define current year, month and location
	year, month, location := m.Month.Current()

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

	return firstOfMonth.AddDate(0, m.Month.Gap, 0)

}

// Get the last day of a previous or next month
func (m *MonthWorker) LastOfLastOfMonth() time.Time {

	// Define current year, month and location
	year, month, location := m.Month.Current()

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

	firstOfLastOfMonth := firstOfMonth.AddDate(0, m.Month.Gap, 0)

	return firstOfLastOfMonth.AddDate(0, int(math.Abs(float64(m.Month.Gap))), -1)

}
