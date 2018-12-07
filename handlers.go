package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"os"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Register method: %s\n", r.Method)

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

		fmt.Println("Created user: ", username)
		checkInternalServerError(err, w)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("login method: %s\n", r.Method)

	if r.Method != "POST" {
		err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	// grab user info from the submitted form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// query database to get match username
	var user User
	err := database.QueryRow("SELECT username, password FROM userInfo WHERE username=?;",
		username).Scan(&user.Username, &user.Password)

	checkInternalServerError(err, w)

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	authenticated = true
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticated = false
	isAuthenticated(w, r)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)

	fmt.Printf("Upload method: %s\n", r.Method)

	if r.Method == "GET" {

		err := tpl.ExecuteTemplate(w, "upload.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	err = r.ParseMultipartForm(32 << 20)

	if err != nil {
		panic(err)
	}

	file, handler, err := r.FormFile("upload-file")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	f, err := os.OpenFile("./posts/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
		return
	}

	defer f.Close()
	io.Copy(f, file)

	var homework Homework
	fmt.Println(homework)

	// Save to database
	//stmt, err := database.Prepare(`
	//	INSERT INTO cost(electric_amount, electric_price, water_amount, water_price, checked_date)
	//	VALUES(?, ?, ?, ?, ?)
	//`)

	//if err != nil {
	//	fmt.Println("Prepare query error")
	//	panic(err)
	//}
	//_, err = stmt.Exec(cost.ElectricAmount, cost.ElectricPrice,
	//	cost.WaterAmount, cost.WaterPrice, cost.CheckedDate)
	//if err != nil {
	//	fmt.Println("Execute query error")
	//	panic(err)
	//}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
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

func postViewHandler(w http.ResponseWriter, _ *http.Request) {

	// TODO: Build this struct based on the information from the database
	hw := Homework{
		Id:        123,
		Title:     "[CS][370][Confer] First Homework",
		PostImage: "image1.jpeg",
		Comments:  []string{"This post is great!", "No, it really isn't"},
	}

	err := tpl.ExecuteTemplate(w, "homework.gohtml", hw)

	checkInternalServerError(err, w)
}
