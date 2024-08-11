package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT NOT NULL,
    done INTEGER NOT NULL,
    created_at DATETIME NOT NULL
);`

func main() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}
}
