package appcrawl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type appStoreReponse struct {
	Count   float64       `json:"resultCount"`
	Results []appStoreApp `json:"results"`
}

type appStoreApp struct {
	Name          string  `json:"trackName"`
	Package       string  `json:"bundleId"`
	PublisherName string  `json:"artistName"`
	StoreUrl      string  `json:"trackViewUrl"`
	IconUrl       string  `json:"artworkUrl60"`
	AppVersion    string  `json:"version"`
	Rating        float32 `json:"averageUserRating"`
	RatingCount   int     `json:"userRatingCount"`
}

type AppStore struct{}

func (a AppStore) searchUrl(sr SearchRequest) string {
	u, _ := url.Parse("https://itunes.apple.com/search")
	q := u.Query()

	if sr.Country != "" {
		q.Set("country", sr.Country)
	}
	if sr.Limit > 0 {
		q.Set("limit", strconv.Itoa(sr.Limit))
	}
	if sr.Query != "" {
		q.Set("term", sr.Query)
	}
	q.Set("media", "software")
	u.RawQuery = q.Encode()
	return u.String()
}

func (a AppStore) Search(sr SearchRequest) ([]App, error) {
	url := a.searchUrl(sr)
	result, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	// TODO: make this steaming, so it can stop after
	// finding a specific bundle

	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var appresults appStoreReponse
	err = json.Unmarshal(response, &appresults)
	if err != nil {
		return nil, err
	}

	apps := make([]App, len(appresults.Results))
	for i, v := range appresults.Results {
		apps[i].Name = v.Name
		apps[i].Package = v.Package
		apps[i].PublisherName = v.PublisherName
		apps[i].StoreUrl = v.StoreUrl
		apps[i].IconUrl = v.IconUrl
		apps[i].AppVersion = v.AppVersion
		apps[i].Rating = v.Rating
		apps[i].RatingCount = v.RatingCount
	}

	return apps, nil
}
