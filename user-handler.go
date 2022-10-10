package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
 
func CreateUser(w http.ResponseWriter, r *http.Request) {
 
    w.Header().Set("Content-Type", "application/json")
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    Database.Create(&user)
    json.NewEncoder(w).Encode(user)
 
}
 
func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var Users  []User
    Database.Find(&Users)
    json.NewEncoder(w).Encode(Users)
}
 
func GetUserById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user  User
    Database.First(&user, mux.Vars(r)["eid"])
    json.NewEncoder(w).Encode(user)
}
 
 
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user  User
    Database.First(&user, mux.Vars(r)["eid"])
    json.NewDecoder(r.Body).Decode(&user)
    Database.Save(&user)
    json.NewEncoder(w).Encode(user)
}
 
 
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user  User
    Database.Delete(&user, mux.Vars(r)["eid"])
    json.NewEncoder(w).Encode("User has been deleted!")
}
 
 
