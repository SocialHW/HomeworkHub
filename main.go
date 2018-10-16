/**
 * Entry point into the HomeworkHub web server.
 *
 */
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})



	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	port := ":3000"

	log.Printf("Server running on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
