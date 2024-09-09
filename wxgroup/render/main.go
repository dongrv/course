package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name  string
	Age   int
	Posts []Post
}

type Post struct {
	ID    int
	Title string
}

func handleUserPage(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Alice",
		Age:  25,
		Posts: []Post{
			{ID: 1, Title: "Hello World"},
			{ID: 2, Title: "Go is Fun"},
		},
	}
	tmpl, err := template.ParseFiles("user_info.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleUserPage)

	handler := loggerMiddleware(router)

	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", handler)
}
