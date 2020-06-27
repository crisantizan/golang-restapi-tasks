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
func AddTaskInFile(task structs.CreateTask, tasks *[]structs.Task) structs.Task {
	// get the last id
	lastID := (*tasks)[len(*tasks)-1].ID

	newTask := structs.Task{
		ID:         lastID + 1,
		CreateTask: task,
	}

	// assign the new task
	*tasks = append(*tasks, newTask)

	// convert data to bytes
	jsonBytes, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}

	// write in json file
	ioutil.WriteFile("data.json", jsonBytes, 0644)

	return newTask
}

// BinarySearch in a array
func BinarySearch(arr []structs.Task, id int, min int, max int) int {
	if max >= min {
		// search slice middle
		middle := (min + max) / 2
		// item
		guess := arr[middle]

		// found
		if guess.ID == id {
			return middle
		}

		// continue with other values
		if guess.ID > id {
			return BinarySearch(arr, id, min, middle-1)
		}

		return BinarySearch(arr, id, middle+1, max)
	}

	return -1
}
