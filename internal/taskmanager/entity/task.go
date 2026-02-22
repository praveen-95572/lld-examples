package entity

import "time"

type TaskStatus string
type TaskPriority string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Done       TaskStatus = "DONE"
)

const (
	Low    TaskPriority = "LOW"
	Medium TaskPriority = "MEDIUM"
	High   TaskPriority = "HIGH"
)

type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     time.Time    `json:"due_date"`
	AssigneeID  string       `json:"assignee_id"`
	Comments    []Comment    `json:"comments"`
	CreatedAt   time.Time    `json:"created_at"`
}
