package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var users []User

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Println("server is now listening to localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		postUser(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "error while encoding users", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Users: %v", users)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users = append(users, user)
	fmt.Fprintf(w, "User %v added", user)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is alive")
}
