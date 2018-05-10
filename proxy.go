package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// func SetupProxy(host string) {
// 	http.Handle("/", newMultipleHostReverseProxy())
// }

func newKansasReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		fmt.Print("Proxying: ")
		fmt.Println(req.URL)
		req.URL.Host = "xhaus.com"
		req.URL.Scheme = "http"
		req.Host = "xhaus.com"
	}
	return &httputil.ReverseProxy{Director: director}
}
