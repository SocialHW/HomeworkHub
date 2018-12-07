package main

import (
	"log"
	"net/http"
)

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	if !authenticated {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}
