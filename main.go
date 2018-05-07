package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Kansas!")

	host := flag.String("host", "localhost", "hostname to use for proxying")
	flag.Parse()

	fmt.Println("Using '" + *host + "' as hostname")

	setupControl(*host)
	setupProxy(*host)

	http.ListenAndServe(":8080", nil)
}
