package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// projectList := &tasklist.ProjectList{}
	// tasklist.CreateProject(db)
	// tasklist.AddProject(db, "test123")
	// projectList.AddProject(db, "test4")
	// projectList.AddProject("test3")
	// projectList.AddTask("test task", 0)
	// projectList.AddTask("test task2", 0)
	// projectList.AddTask("test task3", 1)

	// projectList.SetPriority("high", 0, 1)
	// projectList.SetStatus("done", 0, 0)
	// // projectList.DeleteTask(0, 0)
	// // projectList.DeleteTask(0, 0)
	// // projectList.DeleteProject(2)
	// tasklist.ListProjects(db)
	// projectList.ListTasks(0)
	// // projectList.ListTasks(1)
	// // projectList.ListTasks(2)
}
