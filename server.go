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

type Headers map[string]string

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respondWith(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}

	url, newUrl, err := url.FindOrCreateNewUrl(extractUrl(r))

	if err != nil {
		respondWith(w, http.StatusBadRequest, nil)
		return
	}

	var status int
	if newUrl {
		status = http.StatusCreated
	} else {
		status = http.StatusOk
	}

	shortUrl := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	respondWith(w, status, Headers{
		"Location": shortUrl,
	})
}

func respondWith(
	w http.ResponseWriter,
	status int,
	headers Headers,
) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

fund extractUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func main() {
	http.HandleFunc("/api/shorten", Shorten)
	http.HandleFunc("/r/", Redirector)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
