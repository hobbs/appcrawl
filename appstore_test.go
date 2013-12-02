package appcrawl

import (
	"testing"
)

func TestAppStore(t *testing.T) {
	s := AppStore{}

	req := SearchRequest{
		Query:   "facebook",
		Country: "US",
		Limit:   2,
	}

	apps, err := s.Search(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(apps)

	if len(apps) != req.Limit {
		t.Error("Did not get correct number of results from app store")
	}

	if apps[0].Package != "com.facebook.Facebook" {
		t.Error("Facebook app package does not match")
	}
}
