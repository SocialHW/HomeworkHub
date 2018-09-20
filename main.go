package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static/")))

	log.Println("Server running...")
	http.ListenAndServe(":3000", nil)
}
