package main

type Homework struct {
	Id        uint
	Title     string
	PostImage string
	Comments  []string
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}
