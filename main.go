package main

import (
	"net/http"

	"github.com/crisantizan/golang-restapi-tasks/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks", handler.CreateTask).Methods(http.MethodPost)

	http.ListenAndServe(":3000", r)
}
