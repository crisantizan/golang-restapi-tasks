package structs

// Task a simple task
type Task struct {
	ID int `json:"id,omitempty"`
	CreateTask
}

// CreateTask required data to create a new task
type CreateTask struct {
	Body      string `json:"body,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}
