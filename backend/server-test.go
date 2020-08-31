package main

import (
	"net/http"
)

func main() {
	fmt.Println("Simple Web Server for testing")
	http.Handle("/", http.FileServer(http.Dir("./server-test")))
	http.ListenAndServe(":3000", nil)
}