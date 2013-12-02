package appcrawl

import (
	"testing"
)

func TestAppStoreRank(t *testing.T) {
	store := AppStore{}
	rank, err := AppRank(store, "facebook", "com.facebook.Facebook", "US")

	if err != nil {
		t.Error(err)
	}

	if rank != 1 {
		t.Error("Facebook not ranked first for 'facebook' in US App Store")
	}
}
