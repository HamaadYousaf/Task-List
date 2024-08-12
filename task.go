package tasklist

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type taskItem struct {
	Task     string
	Priority string
	Status   string
	Created  time.Time
}

type projectItem struct {
	Name      string
	TaskItems []taskItem
}

type ProjectList []projectItem

func CreateTask(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS task (
        ID             SERIAL    PRIMARY KEY,
        TASK           TEXT      NOT NULL,
        PRIORITY       TEXT      NOT NULL,
        STATUS         TEXT      NOT NULL,
        CREATED        TEXT      NOT NULL,
        project_id serial references project(id)
    )`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func AddTask(db *sql.DB, task string, projectId int) error {

	query := `SELECT * FROM project WHERE id = $1`
	_, err := db.Query(query, projectId)

	if err != nil {
		log.Fatal(err)
	}

	created := time.Now().Format(time.RFC822)
	taskQuery := `INSERT INTO task (task, priority, status, created, project_id) VALUES ($1, $2, $3, $4, $5)`

	if _, err := db.Exec(taskQuery, task, "low", "new", created, projectId); err != nil {
		log.Fatal(err)
	}

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

func (p *ProjectList) SetStatus(status string, projectIndex int, taskIndex int) error {

	if projectIndex < 0 || projectIndex > len(*p)-1 {
		return errors.New("invalid project index")
	}

	project := &(*p)[projectIndex]

	if taskIndex < 0 || taskIndex > len(project.TaskItems)-1 {
		return errors.New("invalid task index")
	}

	project.TaskItems[taskIndex].Status = strings.ToLower(status)

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

func ListTasks(db *sql.DB, projectIndex int) {
	query := `SELECT id, task, priority, status, created FROM task WHERE project_id=$1`
	rows, err := db.Query(query, projectIndex)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
		},
	}

	var cells [][]*simpletable.Cell

	for rows.Next() {
		var task, priority, status, created string
		var id int

		if err := rows.Scan(&id, &task, &priority, &status, &created); err != nil {
			log.Fatal(err)
		}

		if status == "done" {
			task = green(fmt.Sprintf("\u2705 %s", task))
			priority = green(priority)
			status = green(status)
			created = green(created)
		} else if priority == "high" {
			priority = red(priority)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: task},
			{Text: priority},
			{Text: status},
			{Text: created},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
