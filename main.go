package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Post struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

var (
	posts   = make(map[int]Post)
	nextID  = 1
	postsMu sync.Mutex
)

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
