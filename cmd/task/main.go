package main

import (
	"fmt"

	tasklist "github.com/HamaadYousaf/Task-List"
)

func main() {

	projectList := &tasklist.ProjectList{}

	projectList.AddProject("test")
	projectList.AddProject("test2")
	projectList.AddProject("test3")
	projectList.AddTask("test task", "low", 0)
	projectList.AddTask("test task2", "high", 0)
	projectList.AddTask("test task3", "med", 1)

	fmt.Println((*projectList))
}