package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	userMake("homeworkHub")
}

func userMake(name string) {

	reader := bufio.NewReader(os.Stdin)
	//	var username string
	fmt.Println("Enter Test Username: ")
	username, _ := reader.ReadString('\n')

	//	var email string
	fmt.Println("Enter Test Email: ")
	email, _ := reader.ReadString('\n')

	//	var isadmin string
	fmt.Println("Is User Admin?")
	isadmin, _ := reader.ReadString('\n')

	//	var passwordhash string
	fmt.Println("Enter Password Hash: ")
	passwordhash, _ := reader.ReadString('\n')

	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO userInfo (email,username,isAdmin,passwordHash) VALUES('%s', '%s', '%s', '%s');", email, username, isadmin, passwordhash))
	if err != nil {
		panic(err)
	}
}
