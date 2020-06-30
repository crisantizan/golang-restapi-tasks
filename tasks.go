package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateTask data
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

// Task data
type Task struct {
	ID int `json:"id,omitempty"`
	CreateTask
}

// TaskList properties
type TaskList struct {
	data []*Task
}

// GetTasksFromJSONFile reading file in disk
func GetTasksFromJSONFile() []*Task {
	filename := "data.json"

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	var tasks []*Task

	json.Unmarshal(f, &tasks)

	return tasks
}

// GetOne task
func (t *TaskList) GetOne(id int) (Task, error) {
	index := t.BinarySearch(id, 0, len(t.data)-1)

	if index == -1 {
		return Task{}, errors.New("Task not found")
	}

	return *t.data[index], nil
}

// BinarySearch in tasks
func (t *TaskList) BinarySearch(id int, min int, max int) int {
	if max >= min {
		// search slice middle
		middle := (min + max) / 2
		// item
		guess := t.data[middle]

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

// GetLastID of tasks
func (t *TaskList) GetLastID() int {
	return t.data[len(t.data)-1].ID
}

// AddTaskInFile add new task to json file
func (t *TaskList) AddTaskInFile(task CreateTask) Task {
	// get the last id
	lastID := t.GetLastID()

	newTask := &Task{
		ID:         lastID + 1,
		CreateTask: task,
	}

	// assign the new task
	t.data = append(t.data, newTask)

	// convert data to bytes
	jsonBytes, err := json.Marshal(t.data)

	if err != nil {
		fmt.Println(err)
	}

	// write in json file
	ioutil.WriteFile("data.json", jsonBytes, 0644)

	return *newTask
}

// UpdateTaskInFile and locally
func (t *TaskList) UpdateTaskInFile(id int, taskdata CreateTask) (Task, error) {
	// find index
	index := t.BinarySearch(id, 0, len(t.data)-1)

	newTask := &Task{
		ID:         id,
		CreateTask: taskdata,
	}

	// update task
	t.data[index] = newTask

	// convert data to bytes
	jsonBytes, err := json.Marshal(t.data)

	if err != nil {
		return Task{}, errors.New("Task not found")
	}

	// write in json file
	ioutil.WriteFile("data.json", jsonBytes, 0644)

	return *newTask, nil
}
