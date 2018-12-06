/**
 * Entry point into the HomeworkHub web server.
 *
 */
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var (
	tpl           *template.Template
	authenticated = false
	database      *sql.DB
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	initialize_DB()

	/* Route for index page */
	http.HandleFunc("/", indexHandler)

	// Route for static assets
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	// Route for posts
	http.HandleFunc("/h/", postViewHandler)

	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/logout/", logoutHandler)
	http.HandleFunc("/register/", registerHandler)
	http.HandleFunc("/list/", listHandler)
	http.HandleFunc("/create/", newPost)

	port := ":3000"

	log.Printf("Server running on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func initialize_DB() {
	database, _ = sql.Open("sqlite3", "./homeworkHub.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS userInfo (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS post_info (post_id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, title TEXT, file_path TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS comment_section (post_id INTEGER, username TEXT, comment TEXT)")
	statement.Exec()
}
