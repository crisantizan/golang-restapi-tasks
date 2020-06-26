package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/crisantizan/golang-restapi-tasks/helper"
	"github.com/crisantizan/golang-restapi-tasks/structs"
)

var tasks = helper.ReadFile()

// GetTasks get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// read data sent by client
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	// save here the previous value
	var newTask structs.CreateTask
	// transform to struct
	json.Unmarshal(d, &newTask)

	// validate struct
	if err := newTask.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)

		return
	}

	// save in file and return full task (with id)
	fullNewTask := helper.AddTaskInFile(newTask, &tasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// sent as bytes
	json.NewEncoder(w).Encode(fullNewTask)
}
