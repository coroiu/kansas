package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Welcome to Kansas!")

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /test request")
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
