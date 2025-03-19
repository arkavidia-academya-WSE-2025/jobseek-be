package entity

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Token string `json:"token"`
}
