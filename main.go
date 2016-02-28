// Copyright 2016 Hendrik Will<hendrikwill(at)gmail.com>. All rights reserved.
// License can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	host    string
	port    int
	address string
)

// httpecho is a HTTP server which responds with the requested HTTP status code.
func main() {
	flag.StringVar(&host, "host", "localhost", "Host for HTTP server")
	flag.IntVar(&port, "port", 80, "Port to bind HTTP server to")
	flag.Parse()
	address = fmt.Sprintf("%s:%d", host, port)

	finalHandler := http.HandlerFunc(handler)
	http.Handle("/", addDefaultHeaders(finalHandler))
	log.Fatal(http.ListenAndServe(address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := splitURLPath(r.URL.Path)
	s, err := strconv.Atoi(p[0])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "No status code, Error: %s", err)
	}

	// Redirect for HTTP status codes 301 and 307
	// E.g. curl --location /301/200 is redirected to /200
	if len(p) > 1 &&
		s == http.StatusMovedPermanently || s == http.StatusTemporaryRedirect {
		w.Header().Add("Location", "http://"+address+"/"+p[1])
		http.Redirect(w, r, "http://"+address+"/"+p[1], s)
	}

	// Parameter to specify a delay for the response
	if d, err := strconv.Atoi(r.URL.Query().Get("d")); err == nil {
		time.Sleep(time.Duration(d) * time.Second)
	}

	w.WriteHeader(s)
	fmt.Fprintf(w, "Status Code: %d %s", s, http.StatusText(s))
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

func splitURLPath(path string) []string {
	p := strings.TrimPrefix(path, "/")
	p = strings.TrimSuffix(p, "/")
	return strings.Split(p, "/")
}
