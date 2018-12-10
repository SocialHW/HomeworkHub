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

	type loginPageDataStruct = struct {
		Failed        bool
		Authenticated bool
	}

	if r.Method != "POST" {
		if !authenticated {
			_ = tpl.ExecuteTemplate(w, "login.gohtml", loginPageDataStruct{false, false})

		} else {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		return
	}

	// grab user info from the submitted form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// query database to get match username
	_ = database.QueryRow("SELECT username, password FROM userInfo WHERE username=?;",
		username).Scan(&user.Username, &user.Password)

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err == nil {
		authenticated = true
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		_ = tpl.ExecuteTemplate(w, "login.gohtml", loginPageDataStruct{true, false})
	}

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

	var post Homework

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
	var posts []Homework

	rows, err := database.Query("SELECT * FROM postInfo")

	for rows.Next() {
		var curPost Homework
		if err := rows.Scan(&curPost.Id, &curPost.Username, &curPost.Title, &curPost.Extension); err != nil {
			log.Fatal(err)
		}

		curPost.PostImage = fmt.Sprintf("%d%s", curPost.Id, curPost.Extension)
		posts = append(posts, curPost)
	}

	indexData := struct {
		Authenticated bool
		Posts         []Homework
	}{
		authenticated,
		posts,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", indexData)

	checkInternalServerError(err, w)

}

func postViewHandler(w http.ResponseWriter, r *http.Request) {
	var post Homework

	post.Id, _ = strconv.ParseInt(strings.Replace(r.URL.Path, "/h/", "", 1), 10, 32)

	err := database.QueryRow("SELECT title, username, extension FROM postInfo WHERE postId=?;",
		post.Id).Scan(&post.Title, &post.Username, &post.Extension)

	checkInternalServerError(err, w)

	var comments []string

	rows, err := database.Query("SELECT comment FROM commentSection WHERE postId=?;", post.Id)

	for rows.Next() {
		var curComment string
		if err := rows.Scan(&curComment); err != nil {
			log.Fatal(err)
		}

		comments = append(comments, curComment)
	}

	hw := Homework{
		Id:        post.Id,
		Title:     post.Title,
		PostImage: fmt.Sprintf("%d%s", post.Id, post.Extension),
		Username:  post.Username,
		Comments:  comments,
	}

	postViewData := struct {
		Authenticated bool
		Hw            Homework
	}{
		authenticated,
		hw,
	}

	err = tpl.ExecuteTemplate(w, "homework.gohtml", postViewData)

	checkInternalServerError(err, w)
}

func commentHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(strings.Replace(r.URL.Path, "/comment/", "", 1), 10, 32)

	_, err = database.Exec("INSERT INTO commentSection(postId, comment) VALUES(?, ?);", id, r.FormValue("comment"))

	http.Redirect(w, r, fmt.Sprintf("/h/%d", id), http.StatusMovedPermanently)

}
