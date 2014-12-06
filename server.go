package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vagnerzampieri/shortlink/url"
)

var (
	logOn   *bool
	port    *int
	urlBase string
)

func init() {
	domain := flag.String("d", "localhost", "domain")
	port = flag.Int("p", 4000, "port")
	logOn = flag.Bool("l", true, "log on/off")

	flag.Parse()

	urlBase = fmt.Sprintf("http://%s:%d", *domain, *port)
}

type Headers map[string]string

type Redirector struct {
	stats chan string
}

func (r *Redirector) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	getUrlAndRun(w, req, func(url *url.Url) {
		http.Redirect(w, req, url.Destination, http.StatusMovedPermanently)
		r.stats <- url.Id
	})
}

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
		status = http.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	respondWith(w, status, Headers{
		"Location": shortUrl,
		"Link":     fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", urlBase, url.Id),
	})

	logging("URL %s successfully shortened to %s.", url.Destination, shortUrl)
}

func Viewer(w http.ResponseWriter, r *http.Request) {
	getUrlAndRun(w, r, func(url *url.Url) {
		json, err := json.Marshal(url.Stats())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, string(json))
	})
}

func getUrlAndRun(w http.ResponseWriter, r *http.Request, executor func(*url.Url)) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		executor(url)
	} else {
		http.NotFound(w, r)
	}
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

func respondWithJSON(w http.ResponseWriter, response string) {
	respondWith(w, http.StatusOK, Headers{"Content-Type": "application/json"})
	fmt.Fprintf(w, response)
}

func extractUrl(r *http.Request) string {
	rawBody := make([]byte, r.ContentLength)
	r.Body.Read(rawBody)
	return string(rawBody)
}

func recordStatistics(stats <-chan string) {
	for id := range stats {
		url.RegisterClick(id)
		logging("Click successfully regitered for %s.", id)
	}
}

func logging(format string, vals ...interface{}) {
	if *logOn {
		log.Printf(fmt.Sprintf("%s\n", format), vals...)
	}
}

func main() {
	url.ConfigRepository(url.NewMemoRepository())

	stats := make(chan string)
	defer close(stats)
	go recordStatistics(stats)

	http.Handle("/r/", &Redirector{stats})
	http.HandleFunc("/api/shorten", Shorten)
	http.HandleFunc("/api/stats", Viewer)

	logging("Initializer server in port %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
