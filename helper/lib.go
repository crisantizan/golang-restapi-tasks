package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/crisantizan/golang-restapi-tasks/structs"
)

// ReadFile read a json file
func ReadFile() (tasks []structs.Task) {
	filename := "data.json"

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(f, &tasks)

	return
}

// AddTaskInFile add new task to json file
func AddTaskInFile(task structs.CreateTask, tasks []structs.Task) bool {
	newTask := structs.Task{
		ID:         len(tasks) + 1,
		CreateTask: task,
	}

	data := append(tasks, newTask)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("data.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return true
}
