package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	length  = 5
	symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

type Repository interface {
	IdExist(id string) bool
	FindId(id string) *Url
	FindUrl(url string) *Url
	Save(url Url) error
	RegisterClick(id string)
	FindClicks(id string) int
}

type Url struct {
	Id          string    `json:"id"`
	Criation    time.Time `json:"criation"`
	Destination string    `json:"destination"`
}

type Stats struct {
	Url    *Url `json:"url"`
	Clicks int  `json:"clicks"`
}

var repo Repository

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ConfigRepository(r Repository) {
	repo = r
}
