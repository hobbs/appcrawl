package appcrawl

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "App not found"
}

func AppRank(store Store, keyword string, bundle string, country string) (int, error) {

	req := SearchRequest{
		Query:   keyword,
		Country: country,
		Limit:   200,
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
