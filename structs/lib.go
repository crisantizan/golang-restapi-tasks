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

// Tasks manage
type Tasks struct {
	Data []Task
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

// BinarySearch in a array
func (t Tasks) BinarySearch(id int, min int, max int) int {
	if max >= min {
		// search slice middle
		middle := (min + max) / 2
		// item
		guess := t.Data[middle]

		// found
		if guess.ID == id {
			return middle
		}

		// continue with other values
		if guess.ID > id {
			return t.BinarySearch(id, min, middle-1)
		}

		return t.BinarySearch(id, middle+1, max)
	}

	return -1
}
