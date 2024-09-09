package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := user{Name: "John Doe", Age: 30}
	respondWithJSON(w, http.StatusOK, user)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	mux.HandleFunc("/api/user", userHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.Handle("/", http.FileServer(http.Dir("./public")))
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(mux),
	}

	log.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
