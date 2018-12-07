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
	err           error
	user          User
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	initializeDb()
	defer func() {
		err = database.Close()
		if err != nil {
			panic(err)
		}
	}()

	// Route for index page
	http.HandleFunc("/", indexHandler)

	// Route for static assets
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	// Route for posts
	http.HandleFunc("/h/", postViewHandler)
	http.Handle("/h/img/", http.StripPrefix("/h/img", http.FileServer(http.Dir("./posts"))))

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/upload", uploadHandler)

	port := ":3000"

	log.Printf("Server running on port %s...\n", port)
	serve := http.ListenAndServe(port, nil)

	if serve != nil {
		panic(serve)
	}
}

func initializeDb() {
	database, _ = sql.Open("sqlite3", "./db.sqlite3")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS userInfo (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT)")
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS postInfo (postId INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, title TEXT, extension TEXT)")
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS commentSection (postId INTEGER, username TEXT, comment TEXT)")
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	err = database.Ping()

	if err != nil {
		panic(err)
	}

}
