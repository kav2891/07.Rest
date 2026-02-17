package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"tasks-api/internal/models"
	"tasks-api/internal/storage"
)

type Handler struct {
	Store storage.Storage
}

func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// /tasks
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		tasks := h.Store.List()
		writeJSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		created, err := h.Store.Create(task)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// /tasks/{id}
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	switch r.Method {

	case http.MethodGet:
		task, ok := h.Store.Get(id)
		if !ok {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeJSON(w, http.StatusOK, task)

	case http.MethodPut:
		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		updated, err := h.Store.Update(id, task)
		if err != nil {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}

		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
