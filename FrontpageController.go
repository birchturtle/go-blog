package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Blog World (online, too!)")
}
func fisk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, fish friend %s", r.URL.Query().Get("mig"))
}
