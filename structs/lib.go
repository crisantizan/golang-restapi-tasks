package structs

import (
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

// Validate struct properties
func (t CreateTask) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Body, validation.Required),
	)
}
