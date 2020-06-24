package main

import (
	"github.com/crisantizan/golang-restapi-tasks/helper"
	"github.com/crisantizan/golang-restapi-tasks/structs"
)

func main() {
	helper.AddTaskInFile(structs.CreateTask{
		Body:      "Body",
		Completed: true,
	})
	// f, err := ioutil.ReadFile("./data.json")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var tasks []Task

	// json.Unmarshal(f, &tasks)

	// fmt.Println(tasks)
}
