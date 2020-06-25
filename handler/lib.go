package handler

import (
	"encoding/json"
	"net/http"

	"github.com/crisantizan/golang-restapi-tasks/helper"
)

// GetTasks get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := helper.ReadFile()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tasks)
}
