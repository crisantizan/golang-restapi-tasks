package main

import (
	"github.com/crisantizan/golang-restapi-tasks/helper"
	"github.com/crisantizan/golang-restapi-tasks/structs"
)

func main() {
	tasks := helper.ReadFile()

	// fmt.Println(tasks)
	helper.AddTaskInFile(structs.CreateTask{
		Body:      "Body",
		Completed: true,
	}, tasks)
	// f, err := ioutil.ReadFile("./data.json")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var tasks []Task

	// json.Unmarshal(f, &tasks)

	// fmt.Println(tasks)
}
