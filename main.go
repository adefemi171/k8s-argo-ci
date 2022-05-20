package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello new User 1")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":3000", nil)
}
