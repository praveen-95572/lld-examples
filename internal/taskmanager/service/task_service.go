package service

import (
	"errors"
	"lld-examples/internal/taskmanager/entity"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskService struct {
	tasks map[string]*entity.Task
	mu    sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[string]*entity.Task),
	}
}

func (s *TaskService) CreateTask(task *entity.Task) *entity.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = uuid.NewString()
	task.Status = entity.Todo
	task.CreatedAt = time.Now()

	s.tasks[task.ID] = task
	return task
}

func (s *TaskService) GetTask(id string) (*entity.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) UpdateStatus(id string, status entity.TaskStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	task.Status = status
	return nil
}

func (s *TaskService) AddComment(id, content, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	comment := entity.Comment{
		ID:        uuid.NewString(),
		Content:   content,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	task.Comments = append(task.Comments, comment)
	return nil
}

func (s *TaskService) ListTasks() []*entity.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []*entity.Task
	for _, t := range s.tasks {
		result = append(result, t)
	}
	return result
}
