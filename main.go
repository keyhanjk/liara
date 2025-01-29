package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "goodbye!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
