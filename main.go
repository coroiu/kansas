package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Kansas!")

	host := *flag.String("host", "localhost", "hostname to use for proxying")
	controlEndpoint := "api." + host + "/"
	flag.Parse()

	fmt.Println("Using '" + host + "' as hostname")

	http.Handle(controlEndpoint, NewApiServer())
	http.Handle("/", NewKansasReverseProxy())
	http.ListenAndServe(":8080", nil)
}
