package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if authenticated {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}
}

func registerDataHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// grab user info
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Name entered: %s \tPass entered: %s\n", username, password)

	// Check existence of user
	var user User
	err := database.QueryRow("SELECT username, password",
		username).Scan(&user.Username, &user.Password)

	switch {
	// user is available
	case err == sql.ErrNoRows:
		//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		//checkInternalServerError(err, w)
		// insert to database
		err = database.Ping()
		if err != nil {
			panic(err)
		}
		statement, _ := database.Prepare("INSERT INTO userinfo(username, password) VALUES(?, ?)")
		statement.Exec(username, password)

		fmt.Println("Created user: ", username)
		checkInternalServerError(err, w)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if authenticated {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
		checkInternalServerError(err, w)

		return
	}

	// grab user info from the submitted form
	username := r.FormValue("usrname")
	password := r.FormValue("psw")

	// query database to get match username
	var user User
	err := database.QueryRow("SELECT username, password FROM users WHERE username=?",
		username).Scan(&user.Username, &user.Password)

	checkInternalServerError(err, w)

	// validate password
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil || password != user.Password {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

	authenticated = true
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticated = false
	isAuthenticated(w, r)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}

	//rows, err := database.Query("SELECT * FROM homework")
	//checkInternalServerError(err, w)
	//var funcMap = tpl.FuncMap{
	//	"multiplication": func(n float64, f float64) float64 {
	//		return n * f
	//	},
	//	"addOne": func(n int) int {
	//		return n + 1
	//	},
	//}
	//var homeworks []Homework
	//var homework Homework
	//for rows.Next() {
	//	err = rows.Scan(&homework.Id, &homework.ElectricAmount,
	//		&homework.ElectricPrice, &homework.WaterAmount, &homework.WaterPrice, &homework.CheckedDate)
	//	checkInternalServerError(err, w)
	//	homeworks = append(homeworks, homework)
	//}
	//t, err := tpl.New("list.html").Funcs(funcMap).ParseFiles("templates/list.html")
	//checkInternalServerError(err, w)
	//err = t.Execute(w, homeworks)
	//checkInternalServerError(err, w)

}

func newPost(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
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

	http.Redirect(w, r, "/", 301)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

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

func postViewHandler(w http.ResponseWriter, req *http.Request) {

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

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	if !authenticated {
		http.Redirect(w, r, "/login", 301)
	}
}
