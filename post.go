package main

type Post struct {
	title      string `json: "title"`
	desc       string `json: "desc"`
	photo      string `json: "photo"`
	username   string `json: "username"`
	categories string `json: "categories"`
}
