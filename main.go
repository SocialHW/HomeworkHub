package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var database *sql.DB

func main() {
	intitalize_DB()
	http.Handle("/", http.FileServer(http.Dir("static/")))

	log.Println("Server running...")
	http.ListenAndServe(":3000", nil)
}
func initialize_DB() {
	database, _ = sql.Open("sqlite3", "./homeworkHub.db")
	statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS userInfo (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT)")
	statement1.Exec()
	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS post_info (post_id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, title TEXT, file_path TEXT)")
	statement2.Exec()
	statement3, _ := database.Prepare("CREATE TABLE IF NOT EXISTS comment_section (post_id INTEGER, username TEXT, comment TEXT)")
	statement3.Exec()
}
