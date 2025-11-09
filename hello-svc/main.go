package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		res := r.RequestURI + " | " + uuid.New().String()
		fmt.Println(res)
		fmt.Fprintln(w, res)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
