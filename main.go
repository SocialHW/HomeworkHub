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
	/* Route for index page */
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", struct{ Posts []homework }{
			[]homework{
				{
					Id:        123,
					Title:     "[CS][370][Confer] First Homework",
					PostImage: "image1.jpeg",
					Upvotes:   1,
					Downvotes: 99,
					Comments:  []string{"This post is great!", "No, it really isn't"},
					Tags:      []string{"2018", "MAT", "413", "Andriamanalimanana"},
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

		hw := homework{
			Id:        123,
			Title:     "[CS][370][Confer] First Homework",
			PostImage: "image1.jpeg",
			Upvotes:   1,
			Downvotes: 99,
			Comments:  []string{"This post is great!", "No, it really isn't"},
			Tags:      []string{"2018", "MAT", "413", "Andriamanalimanana"},
		}

		err := tpl.ExecuteTemplate(w, "homework.gohtml", hw)

		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	port := ":3000"

	log.Printf("Server running on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

type homework struct {
	Id        uint
	Title     string
	PostImage string
	Upvotes   uint
	Downvotes uint
	Comments  []string
	Tags      []string
}
