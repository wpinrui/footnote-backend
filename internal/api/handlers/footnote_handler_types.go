package handlers

type CreateFootnoteRequest struct {
	Content string `json:"content"`
	Day     string `json:"day"`
}

type CreateFootnoteResponse struct {
	ID int `json:"id" example:"123"`
}

type UpdateFootnoteRequest struct {
	Content string `json:"content"`
}
