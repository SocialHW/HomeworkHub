package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./homeworkHub.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS post_info (post_id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, title TEXT, file_path TEXT)")
	statement.Exec()
}
