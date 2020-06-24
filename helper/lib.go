package helper

import (
	"fmt"

	"github.com/crisantizan/golang-restapi-tasks/structs"
)

// AddTaskInFile add new task to json file
func AddTaskInFile(task structs.CreateTask) {
	fmt.Println(task)
}
