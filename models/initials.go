package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}
type Bill struct {
	gorm.Model
	CompanyID int
	Company   Company
	Title     string  `json:"title" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	DateID    int
	Date      Date `json:"date" binding:"required"`
	TagID     int
	Tag       Tag
}

type Company struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type Date struct {
	gorm.Model
	TimeStamp time.Time `json:"timeStamp" binding:"required"`
}

// func (b Bill) GetTimeStamp() bool {
// 	fmt.Println(b.Tags[0].Name)
// 	return true
// }
