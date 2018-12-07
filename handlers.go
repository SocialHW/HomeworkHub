package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	// grab user info
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Name entered: %s\tPass entered: %s\n", username, password)

	// Check existence of user
	var user User
	err := database.QueryRow("SELECT username, password FROM userInfo WHERE username=?;",
		username).Scan(&user.Username, &user.Password)

	switch {
	// user is available
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkInternalServerError(err, w)
		// insert to database

		_, err = database.Exec("INSERT INTO userInfo(username, password) VALUES(?, ?);", username, hashedPassword)

		checkInternalServerError(err, w)
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	// grab user info from the submitted form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// query database to get match username
	err := database.QueryRow("SELECT username, password FROM userInfo WHERE username=?;",
		username).Scan(&user.Username, &user.Password)

	checkInternalServerError(err, w)

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	checkInternalServerError(err, w)

	authenticated = true

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticated = false
	isAuthenticated(w, r)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)

	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "upload.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	err = r.ParseMultipartForm(32 << 20)

	checkInternalServerError(err, w)

	var post Post

	post.Username = user.Username

	file, handler, err := r.FormFile("upload-file")

	checkInternalServerError(err, w)

	rows, err := database.Query("SELECT COUNT(*) FROM postInfo")

	checkInternalServerError(err, w)
	defer rows.Close()
	var count int64
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	post.Id = count + 1

	post.Title = r.FormValue("title")

	log.Printf("ID: %d\n", post.Id)

	checkInternalServerError(err, w)

	defer file.Close()

	// Regex to match the file extension
	reg, _ := regexp.Compile("\\.[0-9a-z]{1,5}$")
	post.Extension = string(reg.Find([]byte(handler.Filename)))

	filename := fmt.Sprintf("%d%s", post.Id, post.Extension)

	f, err := os.OpenFile("./posts/"+filename, os.O_WRONLY|os.O_CREATE, 0666)

	checkInternalServerError(err, w)

	defer f.Close()
	_, err = io.Copy(f, file)
	checkInternalServerError(err, w)

	_, err = database.Exec("INSERT INTO postInfo(username, title, extension) VALUES(?, ?, ?);",
		post.Username, post.Title, post.Extension)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {

	// TODO: Query the database to populate this array.
	posts := []Homework{
		{
			Id:        123,
			Title:     "[CS][370][Confer] First Homework",
			PostImage: "image1.jpeg",
			Comments:  []string{"This post is great!", "No, it really isn't"},
		},
	}

	indexData := struct {
		Authenticated bool
		Posts         []Homework
	}{
		authenticated,
		posts,
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", indexData)

	checkInternalServerError(err, w)

}

func postViewHandler(w http.ResponseWriter, r *http.Request) {

	var post Post

	post.Id, _ = strconv.ParseInt(strings.Replace(r.URL.Path, "/h/", "", 1), 10, 32)

	err := database.QueryRow("SELECT title, username, extension FROM postInfo WHERE postId=?;",
		post.Id).Scan(&post.Title, &post.Username, &post.Extension)

	checkInternalServerError(err, w)

	// TODO: Build this struct based on the information from the database
	hw := Homework{
		Title:     post.Title,
		PostImage: fmt.Sprintf("%d%s", post.Id, post.Extension),
		Comments:  []string{"This post is great!", "No, it really isn't"},
	}

	err = tpl.ExecuteTemplate(w, "homework.gohtml", hw)

	checkInternalServerError(err, w)
}
