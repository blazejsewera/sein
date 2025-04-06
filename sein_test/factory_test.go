package sein_test

import (
	"github.com/blazejsewera/sein/endpoint"
	"github.com/blazejsewera/sein/resolver"
	gohttp "net/http"
)

const (
	defaultSearchHome = "https://example.com"
	defaultSearch     = "https://example.com?q={{.Query}}"
)

func newSearchQueryHandler() *endpoint.SearchQueryHandler {
	r := resolver.New(defaultSearchHome, defaultSearch)
	return endpoint.NewSearchQueryHandler(r)
}

func newHTTPClient() *gohttp.Client {
	return &gohttp.Client{
		CheckRedirect: func(req *gohttp.Request, via []*gohttp.Request) error {
			return gohttp.ErrUseLastResponse
		},
	}
}
