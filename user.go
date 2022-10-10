package main

type User struct {
	username   string `json: "username"`
	email      string `json: "email"`
	password   string `json: "password"`
	profilepic string `json: "profilePic"`
}
