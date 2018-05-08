package main

import (
	"fmt"
	"net/http"
)

func SetupControl(host string) {
	http.HandleFunc("control."+host+"/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got 'control.host/' request")
		w.Write([]byte("hello world"))
	})
}
