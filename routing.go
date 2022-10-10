package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
 

func HandlerRouting(){
    r:= mux.NewRouter()

    //------------ Employee 
    r.HandleFunc("/employee", CreateEmployee).Methods("POST")
    r.HandleFunc("/employees", GetEmployees).Methods("GET")
    r.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
    r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
    r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")


    //------------- User
    r.HandleFunc("/Users/Register", CreateUser).Methods("POST")
    r.HandleFunc("/Users/Login", GetUserById).Methods("GET")
    r.HandleFunc("/Users", GetUsers).Methods("GET")
    r.HandleFunc("/Users/{username}", UpdateUser).Methods("PUT")
    r.HandleFunc("/Users/{username}", DeleteUser).Methods("DELETE")


    //------------- Post
    r.HandleFunc("/posts", CreatePost).Methods("POST")
    r.HandleFunc("/posts", GetPosts).Methods("GET")
    r.HandleFunc("/posts/{pid}", GetPostById).Methods("GET")
    r.HandleFunc("/posts/{pid}", UpdatePost).Methods("PUT")
    r.HandleFunc("/posts/{pid}", DeletePost).Methods("DELETE")
    
 
    log.Fatal(http.ListenAndServe(":8080",r))  
}
