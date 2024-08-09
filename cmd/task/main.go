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
	projectList.AddTask("test task", 0)
	projectList.AddTask("test task2", 0)
	projectList.AddTask("test task3", 1)

	projectList.SetPriority("high", 0, 1)
	projectList.Complete(0, 0)
	projectList.DeleteTask(0, 0)
	projectList.DeleteTask(0, 0)
	fmt.Println((*projectList))
}