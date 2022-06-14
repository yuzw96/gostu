package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/**
统计一下访问各个api的总次数
*/

var mu sync.Mutex

var count int

func main() {
	http.HandleFunc("/", countHandler)
	http.HandleFunc("/hello", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	//这里使用锁来保证并发情况下count正确的累加
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path =%q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count =%d\n", count)
	mu.Unlock()
}
