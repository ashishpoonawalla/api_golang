package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
 
func Createpost(w http.ResponseWriter, r *http.Request) {
 
    w.Header().Set("Content-Type", "application/json")
    var post Post
    json.NewDecoder(r.Body).Decode(&post)
    Database.Create(&post)
    json.NewEncoder(w).Encode(post)
 
}
 
func Getposts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var posts  []Post
    Database.Find(&posts)
    json.NewEncoder(w).Encode(posts)
}
 
func GetpostById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var post  Post
    Database.First(&post, mux.Vars(r)["eid"])
    json.NewEncoder(w).Encode(post)
}
 
 
func Updatepost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var post  Post
    Database.First(&post, mux.Vars(r)["eid"])
    json.NewDecoder(r.Body).Decode(&post)
    Database.Save(&post)
    json.NewEncoder(w).Encode(post)
}
 
 
func Deletepost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var post  Post
    Database.Delete(&post, mux.Vars(r)["eid"])
    json.NewEncoder(w).Encode("post has been deleted!")
}
 
 
