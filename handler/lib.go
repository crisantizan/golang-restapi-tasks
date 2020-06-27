package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/crisantizan/golang-restapi-tasks/helper"
	"github.com/crisantizan/golang-restapi-tasks/structs"
	"github.com/gorilla/mux"
)

var tasks = structs.Tasks{
	Data: helper.ReadFile(),
}

// custom http response (in JSON)
func httpR(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(structs.Error{
		Method:    r.Method,
		Status:    status,
		Timestamp: time.Now().UTC(),
		Response:  data,
	})
}

// GetTasks get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	httpR(w, r, http.StatusOK, tasks.Data)
}

// GetTask get one task
func GetTask(w http.ResponseWriter, r *http.Request) {
	// get param id and convert to int (is string per default)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// search, -1 if not found
	index := tasks.BinarySearch(id, 0, len(tasks.Data)-1)

	if index == -1 {
		httpR(w, r, http.StatusNotFound, "Task not found")
		return
	}

	httpR(w, r, http.StatusOK, tasks.Data[index])
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
		httpR(w, r, http.StatusBadRequest, err)
		return
	}

	// save in file and return full task (with id)
	fullNewTask := helper.AddTaskInFile(newTask, &tasks.Data)

	httpR(w, r, http.StatusCreated, fullNewTask)
}
