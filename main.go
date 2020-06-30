package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	helpers := &Helper{}
	r := mux.NewRouter()
	h := &Handler{
		TaskList: &TaskList{
			Data: helpers.ReadFile(),
		},
	}

	// GET
	r.HandleFunc("/tasks", h.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id:[0-9]+}", h.GetTask).Methods(http.MethodGet)

	// POST
	r.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)

	// PUT
	r.HandleFunc("/tasks/{id:[0-9]+}", h.UpdateTask).Methods(http.MethodPut)

	http.ListenAndServe(":3000", r)
}
