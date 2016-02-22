// Copyright 2016 Hendrik Will<hendrikwill(at)gmail.com>. All rights reserved.
// License can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// httpecho is a web server which responds with the requested HTTP status code.
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "No status code")
	}

	w.WriteHeader(s)
	fmt.Fprintf(w, "Status Code: %d %s", s, http.StatusText(s))
}
