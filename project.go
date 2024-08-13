package tasklist

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alexeyco/simpletable"
)

func CreateProject(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS project (
        ID             SERIAL    PRIMARY KEY,
        NAME           TEXT      NOT NULL
    )`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func AddProject(db *sql.DB, name string) {

	query := `INSERT INTO project (name) VALUES ($1)`

	if _, err := db.Exec(query, name); err != nil {
		log.Fatal(err)
	}
}

func DeleteProject(db *sql.DB, projectId int) error {

	query := `DELETE FROM project WHERE id = $1`

	if _, err := db.Exec(query, projectId); err != nil {
		log.Fatal(err)
	}

	return nil
}

func ListProjects(db *sql.DB) {

	query := `SELECT * FROM project`
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Project"},
		},
	}

	var cells [][]*simpletable.Cell

	for rows.Next() {
		var name string
		var id int

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: name},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
