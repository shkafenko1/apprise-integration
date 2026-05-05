package models

type Email struct {
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Receiver string `json:"receiver"`
}
