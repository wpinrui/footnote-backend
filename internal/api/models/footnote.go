package models

type Footnote struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
	Day     string `json:"day"` // format: YYYY-MM-DD
}
