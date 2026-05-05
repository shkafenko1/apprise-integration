package dtos

type EmailToSend struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	URLs  string `json:"urls"`
}
