package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Task a simple task
type Task struct {
	ID int `json:"id,omitempty"`
	CreateTask
}

// CreateTask required data to create a new task
type CreateTask struct {
	Body      string `json:"body,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func main() {
	f, err := ioutil.ReadFile("./data.json")

	if err != nil {
		fmt.Println(err)
	}

	var tasks []Task

	json.Unmarshal(f, &tasks)

	fmt.Println(tasks)
}
