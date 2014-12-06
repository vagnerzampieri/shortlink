package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vagnerzampieri/shortlink/url"
)

var (
	port    int
	urlBase string
)

func init() {
	port = 4000
	urlBase = fmt.Sprintf("http://localhost:%d", port)
}

func main() {
	http.HandleFunc("/api/shorten", Shorten)
	http.HandleFunc("/r/", Redirector)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
