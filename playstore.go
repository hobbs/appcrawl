package appcrawl

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PlayStore struct {
}

func (a PlayStore) searchUrl(sr SearchRequest) (playUrl string, postData string) {

	u, _ := url.Parse("https://play.google.com/store/search")
	q := u.Query()

	if sr.Query != "" {
		q.Set("q", sr.Query)
	}
	q.Set("c", "apps")
	u.RawQuery = q.Encode()
	playUrl = u.String()

	post := url.Values{}
	post.Set("start", "0")
	if sr.Limit > 0 {
		post.Set("num", strconv.Itoa(sr.Limit))
	}
	post.Set("numChildren", "0")
	post.Set("ipf", "1")
	post.Set("xhr", "1")

	postData = post.Encode()

	return
}

func (a PlayStore) Search(sr SearchRequest) ([]App, error) {

	playurl, post := a.searchUrl(sr)

	body := strings.NewReader(post)
	req, err := http.NewRequest("POST", playurl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	client := sr.Client
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return nil, err
	}
	cards := doc.Find(".card")
	apps := make([]App, cards.Length())
	cards.Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("data-docid"); ok {
			apps[i].Package = val
		}
		if val, ok := s.Find(".title").Attr("title"); ok {
			apps[i].Name = val
		}

		if val, ok := s.Find(".subtitle").Attr("title"); ok {
			apps[i].PublisherName = val
		}

		if val, ok := s.Find(".card-click-target").Attr("href"); ok {
			apps[i].StoreUrl = "https://play.google.com" + val
		}

		if val, ok := s.Find(".cover-image").Attr("src"); ok {
			apps[i].IconUrl = val
		}

		if val, ok := s.Find(".current-rating").Attr("style"); ok {
			rating := strings.Split(val, " ")
			if len(rating) >= 2 {
				if rating_perc, err := strconv.ParseFloat(strings.TrimRight(rating[1], "%;"), 32); err == nil {
					apps[i].Rating = float32(rating_perc * 5.0 / 100.0)
				}

			}

		}

	})

	return apps, nil
}
