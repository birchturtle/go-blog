package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Blog World (online, too!)")
}

func main() {
	fmt.Println("Hello, Blog World!")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
