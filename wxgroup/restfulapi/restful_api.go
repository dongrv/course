package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = map[int]User{
	1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
	2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var newUser User
	if err := json.Unmarshal(body, &newUser); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	maxId := 0
	for id := range users {
		if id > maxId {
			maxId = id
		}
	}
	newUser.ID = maxId + 1
	users[newUser.ID] = newUser
	json.NewEncoder(w).Encode(newUser)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, ok := users[userId]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var updatedUser User
	if err := json.Unmarshal(body, &updatedUser); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if user, ok := users[userId]; ok {
		user.Name = updatedUser.Name
		user.Email = updatedUser.Email
		users[userId] = user
		json.NewEncoder(w).Encode(user)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if _, ok := users[userId]; ok {
		delete(users, userId)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", handleGetAllUsers).Methods("GET")
	router.HandleFunc("/users", handleCreateUser).Methods("POST")
	router.HandleFunc("/users/{userId}", handleGetUser).Methods("GET")
	router.HandleFunc("/users/{userId}", handleUpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", handleDeleteUser).Methods("DELETE")

	handler := loggerMiddleware(router)

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", handler)
}
