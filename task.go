package tasklist

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/alexeyco/simpletable"
)

func CreateTask(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS task (
        ID             SERIAL    PRIMARY KEY,
        TASK           TEXT      NOT NULL,
        PRIORITY       TEXT      NOT NULL,
        STATUS         TEXT      NOT NULL,
        CREATED        TEXT      NOT NULL,
        project_id serial references project(id)
        ON DELETE CASCADE
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

func SetPriority(db *sql.DB, priority string, taskId int) error {
	updateQuery := `UPDATE task SET priority = $2 WHERE id = $1`

	if _, err := db.Exec(updateQuery, taskId, priority); err != nil {
		log.Fatal(err)
	}

	return nil
}

func SetStatus(db *sql.DB, status string, taskId int) error {

	updateQuery := `UPDATE task SET status = $2 WHERE id = $1`

	if _, err := db.Exec(updateQuery, taskId, status); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteTask(db *sql.DB, taskId int) error {

	query := `DELETE FROM task WHERE id = $1`

	if _, err := db.Exec(query, taskId); err != nil {
		log.Fatal(err)
	}

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
