package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func SetupProxy(host string) {
	http.Handle("test."+host+"/", newMultipleHostReverseProxy())
	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	// handle error
	// }
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		// handle error
	// 	}
	// 	go handleConnection(conn)
	// }

	// http.HandleFunc("test."+host+"/", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Println("got 'test.host/' request")

	// 	// hj, ok := w.(http.Hijacker)
	// 	// if !ok {
	// 	// 	http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
	// 	// 	return
	// 	// }

	// 	// _, bufrw, err := hj.Hijack()
	// 	// if err != nil {
	// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// 	return
	// 	// }

	// 	// tReader, tWriter := connectTo()
	// })
}

func newMultipleHostReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		// dump, _ := httputil.DumpRequest(req, false)
		// fmt.Println(string(dump))
		fmt.Print("Proxying: ")
		fmt.Println(req.URL)
		req.URL.Host = "xhaus.com"
		req.URL.Scheme = "http"
		req.Host = "xhaus.com"
		// target := targets[rand.Int()%len(targets)]
		// req.URL.Scheme = target.Scheme
		// req.URL.Host = target.Host
		// req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

// func handleConnection(conn net.Conn) {
// 	r := bufio.NewReader(conn)
// 	// s := bufio.NewScanner(conn)
// 	h := readHeader(r)
// 	req, _ := http.ReadRequest(h)
// 	fmt.Println(req.Host)
// 	conn.Write([]byte("hej"))
// 	conn.Close()
// }

// func readHeader(reader *bufio.Reader) bufio.Reader {
// 	c := make([]byte, 0, 2048)
// 	for {
// 		p, _ := reader.ReadBytes(4)
// 		if bytes.Equal(p[:4], []byte("\r\n\r\n")) {
// 			fmt.Println(string(c))
// 			return bufio.NewReader(bytes.NewReader(c))
// 		}
// 		c = append(c, p...)
// 	}
// }

func connectTo() (*bufio.Reader, *bufio.Writer) {
	con, err := net.Dial("tcp", "google.se:80")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(con)
	writer := bufio.NewWriter(con)

	return reader, writer
	// tcpcon, ok := con.(*net.TCPConn)

	// if !ok {
	// 	log.Fatal(ok)
	// }

	// return tcpcon
	// } else {
	// 	return
	// }
	// _, err = io.Copy(os.Stdout, con)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// err = con.Close()
	// if err != nil {
	//     log.Fatal(err)
	// }
}
