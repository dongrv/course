package main

import (
	"log"
	"net/http"
)

func main() {
	listen := ":8080"
	dir := "./wasm/html" // 相对curse根目录
	log.Printf("listening on %q...", ":8080")
	err := http.ListenAndServe(listen, http.FileServer(http.Dir(dir)))
	log.Fatalln(err)
}
