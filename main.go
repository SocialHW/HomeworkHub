/**
 * Entry point into the HomeworkHub web server.
 *
 */
package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
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
	http.Handle("/", http.FileServer(http.Dir("static/")))
	/* Route for index page */
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", struct{ Posts []Homework }{
			[]Homework{
				{
					Id:        123,
					Title:     "[CS][370][Confer] First Homework",
					PostImage: "image1.jpeg",
					Comments:  []string{"This post is great!", "No, it really isn't"},
				},
			},
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Route for static assets
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	/* Route for posts */
	http.HandleFunc("/h/", func(w http.ResponseWriter, req *http.Request) {

		hw := Homework{
			Id:        123,
			Title:     "[CS][370][Confer] First Homework",
			PostImage: "image1.jpeg",
			Comments:  []string{"This post is great!", "No, it really isn't"},
		}

		err := tpl.ExecuteTemplate(w, "homework.gohtml", hw)

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/create", createHandler)

	port := ":3000"

	log.Printf("Server running on port %s...\n", port)
	http.ListenAndServe(port, nil)
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
