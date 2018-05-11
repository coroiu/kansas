package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Kansas!")

	initStore()

	host := *flag.String("host", "localhost", "hostname to use for proxying")
	controlEndpoint := "control." + host + "/"
	apiEndpoint := "api." + controlEndpoint
	flag.Parse()

	fmt.Println("Using '" + host + "' as hostname")

	loadConfiguration()
	http.Handle(apiEndpoint, newAPIServer())
	http.Handle("/", newKansasReverseProxy())
	http.ListenAndServe(":8080", nil)
}
