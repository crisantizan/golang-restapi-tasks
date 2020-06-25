package main

import (
	"net/http"

	"github.com/crisantizan/golang-restapi-tasks/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.GetTasks)

	http.ListenAndServe(":3000", r)
}
