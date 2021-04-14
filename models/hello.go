package models

// Message : struct for message
type Message struct {
	Text string `json:"text"`
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
