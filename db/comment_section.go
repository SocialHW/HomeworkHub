package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./db.sqlite3")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS comment_section (post_id INTEGER, username TEXT, comment TEXT)")
	statement.Exec()
}
