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
