package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./homeworkHub.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tags (post_id INTEGER, tag TEXT)")
	statement.Exec()
}
