package appcrawl

import "net/http"

type Store interface {
	Search(req SearchRequest) ([]App, error)
}

type App struct {
	Name          string
	Package       string //aka bundle id
	PublisherName string
	StoreUrl      string
	IconUrl       string
	AppVersion    string
	Rating        float32
	RatingCount   int
}

type SearchRequest struct {
	Query   string
	Country string
	Limit   int
	Client  *http.Client
}
