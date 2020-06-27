package main

import (
	"net/http"

	"github.com/crisantizan/golang-restapi-tasks/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// GET
	r.HandleFunc("/tasks", handler.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/task/{id:[0-9]+}", handler.GetTask).Methods(http.MethodGet)

	// POST
	r.HandleFunc("/tasks", handler.CreateTask).Methods(http.MethodPost)

	http.ListenAndServe(":3000", r)
}
