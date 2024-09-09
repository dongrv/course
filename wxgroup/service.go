package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	db map[string]string
)

func init() {
	db = make(map[string]string)
	for i := 0; i < 100000; i++ {
		db[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
}

func main() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("Could not start pprof server: %v", err)
		}
	}()

	http.HandleFunc("/", handler)
	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing key parameter"))
		return
	}

	// 模拟锁争用
	mu.Lock()
	value, ok := db[key]
	mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key not found"))
		return
	}

	time.Sleep(time.Second) // 模拟 I/O 操作延迟
	w.Write([]byte(value))
}

var rwmu sync.RWMutex

func handler2(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing key parameter"))
		return
	}

	rwmu.RLock()
	value, ok := db[key]
	rwmu.RUnlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key not found"))
		return
	}

	time.Sleep(time.Second) // 模拟 I/O 操作延迟
	w.Write([]byte(value))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing key parameter"))
		return
	}

	go func() {
		time.Sleep(time.Second) // 模拟 I/O 操作延迟
		w.Write([]byte(db[key]))
	}()
}
