// Copyright 2016 Hendrik Will<hendrikwill(at)gmail.com>. All rights reserved.
// License can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// httpecho is a web server which responds with the requested HTTP status code.
func main() {
	clPort := flag.Int("port", 80, "Port to bind server to")
	flag.Parse()

	finalHandler := http.HandlerFunc(handler)
	http.Handle("/", addDefaultHeaders(finalHandler))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*clPort), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "No status code")
	} else {
		w.WriteHeader(s)
		fmt.Fprintf(w, "Status Code: %d %s", s, http.StatusText(s))
	}
}

func addDefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
