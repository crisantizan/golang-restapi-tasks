package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	h := GetHandlers()

	// GET
	r.HandleFunc("/tasks", h.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id:[0-9]+}", h.GetTask).Methods(http.MethodGet)

	// POST
	r.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)

	// PUT
	r.HandleFunc("/tasks/{id:[0-9]+}", h.UpdateTask).Methods(http.MethodPut)

	// DELETE
	r.HandleFunc("/tasks/{id:[0-9]+}", h.DeleteTask).Methods(http.MethodDelete)

	// MIDDLEWARES
	r.Use(LogginMiddleware)

	http.ListenAndServe(":3000", r)
}
