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

func RegisterClick(id string) {
	repo.RegisterClick(id)
}

func FindOrCreateNewUrl(destination string) (u *Url, newUrl bool, err error) {
	if u = repo.FindUrl(destination); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destination); err != nil {
		return nil, false, err
	}

	url := Url{generateId(), time.Now(), destination}
	repo.Save(url)
	return &url, true, nil
}
