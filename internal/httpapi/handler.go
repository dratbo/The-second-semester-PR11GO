package httpapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"example.com/pz11-graphql/internal/service"
)

type Handler struct {
	service *service.TaskService
}

func NewHandler(svc *service.TaskService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.service.ListTasks(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	t, err := h.service.GetTaskByID(r.Context(), id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(t)
}

func (h *Handler) PatchTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		ID          string  `json:"id"`
		Title       string  `json:"title"`
		Description *string `json:"description"`
		Done        *bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		id = body.ID
	}

	title := &body.Title
	if body.Title == "" {
		title = nil
	}

	if _, err := h.service.UpdateTask(r.Context(), id, title, body.Description, body.Done); err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
