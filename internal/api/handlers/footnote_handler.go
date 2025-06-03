package handlers

import (
	"encoding/json"
	"footnote-backend/internal/api/middleware"
	"footnote-backend/internal/api/models"
	"footnote-backend/internal/db/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FootnoteHandler struct {
	FootnoteRepository *repositories.FootnoteRepository
}

func NewFootnoteHandler(fr *repositories.FootnoteRepository) *FootnoteHandler {
	return &FootnoteHandler{
		FootnoteRepository: fr,
	}
}

func (fh *FootnoteHandler) CreateFootnote(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateFootnoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	footnote := &models.Footnote{
		UserId:  userId,
		Content: req.Content,
	}

	id, err := fh.FootnoteRepository.Create(footnote)
	if err != nil {
		http.Error(w, "Failed to create footnote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (fh *FootnoteHandler) GetFootnotes(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	footnotes, err := fh.FootnoteRepository.ListByUser(userId)
	if err != nil {
		// Log the error before responding
		// You can use the standard log package or a structured logger if available
		// Example with standard log:
		log.Printf("failed to get footnotes for user %d: %v", userId, err)
		http.Error(w, "Failed to get footnotes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(footnotes)
}

func (fh *FootnoteHandler) GetFootnoteByID(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	footnote, err := fh.FootnoteRepository.GetByID(id, userId)
	if err != nil || footnote == nil {
		http.Error(w, "Footnote not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(footnote)
}

func (fh *FootnoteHandler) UpdateFootnote(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req UpdateFootnoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = fh.FootnoteRepository.Update(id, userId, req.Content)
	if err != nil {
		http.Error(w, "Failed to update footnote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (fh *FootnoteHandler) DeleteFootnote(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = fh.FootnoteRepository.Delete(id, userId)
	if err != nil {
		http.Error(w, "Failed to delete footnote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (fh *FootnoteHandler) SearchFootnotes(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.UserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing search query", http.StatusBadRequest)
		return
	}

	results, err := fh.FootnoteRepository.Search(userId, query)
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(results)
}
