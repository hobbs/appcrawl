package appcrawl

import (
	"strings"
	"testing"
)

func TestPlayStore(t *testing.T) {
	s := PlayStore{}

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

	if apps[0].Package != "com.facebook.katana" {
		t.Error("Facebook app package does not match")
	}

	if !strings.Contains(apps[0].PublisherName, "Facebook") {
		t.Error("Facebook not publisher")
	}

	if !strings.Contains(apps[0].Name, "Facebook") {
		t.Error("Facebook not in app name")
	}

	if len(apps[0].StoreUrl) == 0 {
		t.Error("Missing store URL")
	}

	if len(apps[0].IconUrl) == 0 {
		t.Error("Missing icon URL")
	}

	if apps[0].Rating <= 0 || apps[0].Rating > 5.0 {
		t.Error("Invalid app rating")
	}
}
