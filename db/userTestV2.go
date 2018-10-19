package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

func main() {
	userMake("homeworkHubUser")
}

func userMake(name string) {

	reader := bufio.NewReader(os.Stdin)
	//	var username string
	fmt.Println("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.Replace(username, "\n", "", -1)
	//	var email string
	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)

	//	var passwordhash string
	fmt.Println("Enter Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", "root:password@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO userInfo (email,username,passwordHash) VALUES('%s', '%s', '%s');", email, username, hash))
	if err != nil {
		panic(err)
	}
}
