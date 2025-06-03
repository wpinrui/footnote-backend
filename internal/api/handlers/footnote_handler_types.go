package handlers

type CreateFootnoteRequest struct {
	Content string `json:"content"`
}

type UpdateFootnoteRequest = CreateFootnoteRequest
