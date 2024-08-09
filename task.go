package tasklist

import (
	"errors"
	"strings"
)

type taskItem struct {
	Task     string
	Priority string
	Done     bool
}

type projectItem struct {
	Name string
	TaskItems []taskItem
}

type ProjectList []projectItem

func (p *ProjectList) AddProject(name string) {

	newProject := projectItem{
		Name:     name,
		TaskItems: make([]taskItem, 0, 1),
	}

	*p = append(*p, newProject)
}

func (p *ProjectList) AddTask(task string, projectIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid index")
	}

	newTask := taskItem{
		Task:     task,
		Priority: "normal",
		Done:     false,
	}

	project := &(*p)[projectIndex]
	project.TaskItems = append(project.TaskItems, newTask)
	
	return nil
}

func (p *ProjectList) SetPriority(priority string, projectIndex int, taskIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid project index")
	}

	project := &(*p)[projectIndex]

	if taskIndex < 0 || taskIndex > len(project.TaskItems)-1 {
		return errors.New("invalid task index")
	}

	project.TaskItems[taskIndex].Priority = strings.ToLower(priority)
	
	return nil
}

func  (p *ProjectList) Complete(projectIndex int, taskIndex int) error{

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid project index")
	}

	project := &(*p)[projectIndex]

	if taskIndex < 0 || taskIndex > len(project.TaskItems)-1 {
		return errors.New("invalid task index")
	}

	project.TaskItems[taskIndex].Done = true

	return nil
}

func (p *ProjectList) DeleteTask(projectIndex int, taskIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid project index")
	}

	project := &(*p)[projectIndex]

	if taskIndex < 0 || taskIndex > len(project.TaskItems)-1 {
		return errors.New("invalid task index")
	}

	project.TaskItems = append(project.TaskItems[:taskIndex], project.TaskItems[taskIndex+1:]...)

	return nil
}