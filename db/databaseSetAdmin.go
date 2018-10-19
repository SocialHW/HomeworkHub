package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	set("homeworkHubUser")
}

func set(name string) {
	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("ALTER TABLE userInfo ALTER isAdmin SET DEFAULT false;")
	if err != nil {
		panic(err)
	}
}
