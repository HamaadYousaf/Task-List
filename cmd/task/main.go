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

	tasklist.ListProjects(db)
}
