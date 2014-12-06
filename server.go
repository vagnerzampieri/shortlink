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
	urlBase = fmt.Sprintf(":%d", port)
}
