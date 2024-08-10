package main

import (
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
	projectList.SetStatus("done", 0, 0)
	// projectList.DeleteTask(0, 0)
	// projectList.DeleteTask(0, 0)
	// projectList.DeleteProject(2)
	projectList.ListProjects()
	projectList.ListTasks(0)
	// projectList.ListTasks(1)
	// projectList.ListTasks(2)
}
