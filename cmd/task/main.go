package main

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

	tasklist.CreateProject(db)
	tasklist.CreateTask(db)

	addProject := flag.Bool("add-project", false, "add a new project")
	delProject := flag.Int("del-project", -1, "delete a project")
	selProject := flag.Int("project", -1, "select a project")
	listProjects := flag.Bool("list-projects", false, "list projects")

	addTask := flag.Bool("add-task", false, "add a new task")
	setPriority := flag.String("priority", "", "set priority")
	setStatus := flag.String("status", "", "set status")
	delTask := flag.Int("del-task", -1, "delete a task")
	listTask := flag.Bool("list-tasks", false, "list tasks")
	selTask := flag.Int("task", -1, "select a task")

	flag.Parse()

	switch {
	case *addProject:
		project, err := getInput(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		tasklist.AddProject(db, project)

	case *delProject >= 0:
		tasklist.DeleteProject(db, *delProject)

	case *listProjects:
		tasklist.ListProjects(db)

	case *addTask:

		if *selProject == -1 {
			fmt.Println("No project specified")
			os.Exit(0)
		}

		task, err := getInput(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		tasklist.AddTask(db, task, *selProject)

	case *setPriority != "":

		if *selTask == -1 {
			fmt.Println("No task specified")
			os.Exit(0)
		}

		tasklist.SetPriority(db, *setPriority, *selTask)

	case *setStatus != "":

		if *selTask == -1 {
			fmt.Println("No task specified")
			os.Exit(0)
		}

		tasklist.SetStatus(db, *setStatus, *selTask)

	case *delTask >= 0:
		tasklist.DeleteTask(db, *delTask)

	case *listTask:
		if *selProject == -1 {
			fmt.Println("No project specified")
			os.Exit(0)
		}

		tasklist.ListTasks(db, *selProject)

	default:
		fmt.Println("invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty input")
	}

	return text, nil
}
