package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	create("homeworkHub")
}

func create(name string) {
	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE userInfo(user_id integer NOT NULL AUTO_INCREMENT PRIMARY KEY, email varchar(32) NOT NULL, username varchar(32) NOT NULL, isAdmin varchar(6), passwordHash varchar(32) NOT NULL  )")
	if err != nil {
		panic(err)
	}
}
