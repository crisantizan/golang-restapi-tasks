package structs

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Task a simple task
type Task struct {
	ID int `json:"id,omitempty"`
	CreateTask
}

// CreateTask required data to create a new task
type CreateTask struct {
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

// Error is a custom HTTP error
type Error struct {
	Method    string      `json:"method,omitempty"`
	Status    int         `json:"status,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Response  interface{} `json:"response,omitempty"`
}

// Validate struct properties
func (t CreateTask) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Body, validation.Required),
	)
}
