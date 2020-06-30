package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Helper struct
type Helper struct{}

// ReadFile read a json file
func (Helper) ReadFile() []*Task {
	filename := "data.json"

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	var tasks []*Task

	json.Unmarshal(f, &tasks)

	return tasks
}
