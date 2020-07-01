package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	TaskList *TaskList
}

// GetHandlers and tasks data
func GetHandlers() *Handler {
	return &Handler{
		TaskList: &TaskList{
			data: GetTasksFromJSONFile(),
		},
	}
}

// Error is a custom HTTP error
type Error struct {
	Method    string      `json:"method,omitempty"`
	Status    int         `json:"status,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Response  interface{} `json:"response,omitempty"`
}

// HTTPResponse json to client
func (Handler) HTTPResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Error{
		Method:    r.Method,
		Status:    status,
		Timestamp: time.Now().UTC(),
		Response:  data,
	})
}

// Redirect to /tasks
func (Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tasks", http.StatusMovedPermanently)
}

// GetTasks handler
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	h.HTTPResponse(w, r, http.StatusOK, h.TaskList.data)
}

// GetTask handler
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	// get param id and convert to int (is string per default)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	task, err := h.TaskList.GetOne(id)

	// not found
	if err != nil {
		h.HTTPResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	h.HTTPResponse(w, r, http.StatusOK, task)
}

// CreateTask handler
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// read data sent by client
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	// save here the previous value
	var newTask CreateTask
	// transform to struct
	json.Unmarshal(d, &newTask)

	// validate struct
	if err := newTask.Validate(); err != nil {
		fmt.Println(err)
		h.HTTPResponse(w, r, http.StatusBadRequest, err)
		return
	}

	// save in file and return full task (with id)
	fullNewTask := h.TaskList.AddTaskInFile(newTask)

	h.HTTPResponse(w, r, http.StatusCreated, fullNewTask)
}

// UpdateTask handler
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		h.HTTPResponse(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	// read data sent by client
	var taskdata CreateTask
	json.Unmarshal(body, &taskdata)

	// validate properties
	if err := taskdata.Validate(); err != nil {
		h.HTTPResponse(w, r, http.StatusBadRequest, err)
		return
	}

	// get param id and convert to int (is string per default)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	taskUpdated, err := h.TaskList.UpdateTaskInFile(id, taskdata)

	if err != nil {
		h.HTTPResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	h.HTTPResponse(w, r, http.StatusOK, taskUpdated)
}

// DeleteTask handler
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.TaskList.DeleteTaskInFile(id); err != nil {
		h.HTTPResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	h.HTTPResponse(w, r, http.StatusOK, true)
}
