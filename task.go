package tasklist

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
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

func (p *ProjectList) DeleteProject(projectIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid project index")
	}

	*p = append((*p)[:projectIndex], (*p)[projectIndex+1:]...)

	return nil
}

func (p *ProjectList) ListProjects() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Project"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, project := range *p {
		name := project.Name
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: name},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (p *ProjectList) ListTasks(projectIndex int) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
		},
	}

	var cells [][]*simpletable.Cell

	tasks := (*p)[projectIndex].TaskItems
	for idx, item := range tasks {
		task := item.Task
		priority := item.Priority
		done := "no"
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			priority = green(item.Priority)
			done = green("yes")
		} else if item.Priority == "high"{
			priority = red(item.Priority)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: priority},
			{Text: done},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
	