package main

type Homework struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	Username  string   `json:"username"`
	PostImage string   `json:"postImage"`
	Extension string   `json:"extension"`
	Comments  []string `json:"comments"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
