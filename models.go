package main

type Homework struct {
	Id        uint
	Title     string
	PostImage string
	Comments  []string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Extension string `json:"title"`
}
