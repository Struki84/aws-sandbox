package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go app is running and serving HTTP request...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":80", nil)
}
