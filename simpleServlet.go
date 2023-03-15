package main

import (
	"fmt"
	"net/http"
)

func main() {
	num := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num = num + 1
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Sprintln(num)
	})

	http.ListenAndServe(":80", nil)
}
