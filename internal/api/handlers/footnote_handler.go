package handlers

import (
	"encoding/json"
	"footnote-backend/internal/api/middleware"
	"footnote-backend/internal/api/models"
	"footnote-backend/internal/db/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"

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

// CreateFootnote godoc
// @Summary Create a new footnote
// @Description Adds a footnote for the authenticated user with content and day
// @Tags footnote
// @Accept json
// @Produce json
// @Param request body CreateFootnoteRequest true "Footnote content and day"
// @Success 201 {object} map[string]int "Returns created footnote ID"
// @Failure 400 {string} string "Footnote content or day cannot be empty / Invalid request payload"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to create footnote"
// @Security ApiKeyAuth
// @Router /footnote [post]
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

	if strings.TrimSpace(req.Content) == "" {
		http.Error(w, "Footnote content cannot be empty", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Day) == "" {
		http.Error(w, "Footnote day cannot be empty", http.StatusBadRequest)
		return
	}

	footnote := &models.Footnote{
		UserId:  userId,
		Content: req.Content,
		Day:     req.Day,
	}

	id, err := fh.FootnoteRepository.Create(footnote)
	if err != nil {
		http.Error(w, "Failed to create footnote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// GetFootnotes godoc
// @Summary Get all footnotes for the authenticated user
// @Description Retrieve a list of footnotes created by the authenticated user
// @Tags footnote
// @Produce json
// @Success 200 {array} models.Footnote
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to get footnotes"
// @Security ApiKeyAuth
// @Router /footnote [get]
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

// GetFootnoteByID godoc
// @Summary Get a footnote by ID
// @Description Retrieves a footnote by its ID for the authenticated user
// @Tags footnote
// @Produce json
// @Param id path int true "Footnote ID"
// @Success 200 {object} models.Footnote
// @Failure 400 {string} string "Invalid ID"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Footnote not found"
// @Security ApiKeyAuth
// @Router /footnote/{id} [get]
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

// UpdateFootnote godoc
// @Summary Update a footnote by ID
// @Description Updates the content of a footnote for the authenticated user
// @Tags footnote
// @Accept json
// @Param id path int true "Footnote ID"
// @Param request body UpdateFootnoteRequest true "Updated footnote content"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid ID or request payload"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to update footnote"
// @Security ApiKeyAuth
// @Router /footnote/{id} [put]
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

// DeleteFootnote godoc
// @Summary Delete a footnote by ID
// @Description Deletes a footnote for the authenticated user
// @Tags footnote
// @Param id path int true "Footnote ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid ID"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to delete footnote"
// @Security ApiKeyAuth
// @Router /footnote/{id} [delete]
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

// SearchFootnotes godoc
// @Summary Search footnotes
// @Description Search footnotes for the authenticated user by query string
// @Tags footnote
// @Param q query string true "Search query"
// @Success 200 {array} models.Footnote
// @Failure 400 {string} string "Missing search query"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Search failed"
// @Security ApiKeyAuth
// @Router /footnote/search [get]
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
