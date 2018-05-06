package main

import "fmt"
import "flag"
import "net/http"

func main() {
	fmt.Println("Welcome to Kansas!")

	host := flag.String("host", "localhost", "hostname to use for proxying")
	flag.Parse()

	fmt.Println("Using '" + *host + "' as hostname")

	http.HandleFunc(*host+"/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /test request")
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
