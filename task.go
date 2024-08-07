package tasklist

import (
	"errors"
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

func (p *ProjectList) AddTask(task string, priority string, projectIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid index")
	}

	newTask := taskItem{
		Task:     task,
		Priority: priority,
		Done:     false,
	}

	project := &(*p)[projectIndex]
	project.TaskItems = append(project.TaskItems, newTask)
	
	return nil
}