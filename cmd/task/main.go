package main

import (
	"database/sql"
	"log"
	"os"

	tasklist "github.com/HamaadYousaf/Task-List"
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

	// tasklist.CreateTask(db)
	// tasklist.AddTask(db, "testtask1", 2)
	// tasklist.AddTask(db, "testtask2", 2)
	// tasklist.AddTask(db, "testtask3", 2)
	// tasklist.AddTask(db, "testtask123", 1)
	// tasklist.AddTask(db, "testtask321", 1)
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
	tasklist.ListTasks(db, 2)
}
