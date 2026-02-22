package handler

import (
	"encoding/json"
	"lld-examples/internal/taskmanager/entity"
	"lld-examples/internal/taskmanager/service"
	"lld-examples/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{Service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	json.NewDecoder(r.Body).Decode(&task)

	created := h.Service.CreateTask(&task)
	utils.JSONResponse(w, created, http.StatusCreated)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := h.Service.GetTask(id)
	if err != nil {
		utils.JSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, task, http.StatusOK)
}

func (h *TaskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload struct {
		Status entity.TaskStatus `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	err := h.Service.UpdateStatus(id, payload.Status)
	if err != nil {
		utils.JSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, "status updated", http.StatusOK)
}

func (h *TaskHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload struct {
		Content string `json:"content"`
		UserID  string `json:"user_id"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	err := h.Service.AddComment(id, payload.Content, payload.UserID)
	if err != nil {
		utils.JSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, "comment added", http.StatusOK)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Service.ListTasks()
	utils.JSONResponse(w, tasks, http.StatusOK)
}
