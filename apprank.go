package appcrawl

import "net/http"

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "App not found"
}

func AppRank(store Store, keyword string, bundle string, country string, client *http.Client) (int, error) {

	req := SearchRequest{
		Query:   keyword,
		Country: country,
		Limit:   200,
		Client:  client,
	}

	apps, err := store.Search(req)

	if err != nil {
		return 0, err
	}

	for i, v := range apps {
		if v.Package == bundle {
			return i + 1, nil
		}
	}

	return 0, &NotFoundError{}
}
