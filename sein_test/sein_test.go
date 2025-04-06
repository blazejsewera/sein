package sein_test

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRedirectToDefaultSearchSite(t *testing.T) {
	s := httptest.NewServer(newSearchQueryHandler())
	defer s.Close()
	client := newHTTPClient()

	t.Run("when there is no query", func(t *testing.T) {
		query := ""
		expectedLocation := "https://example.com"
		testRequest(t, client, s.URL, query, expectedLocation)
	})

	t.Run("when the query doesn't include any special command", func(t *testing.T) {
		query := "abc"
		expectedLocation := "https://example.com?q=" + url.QueryEscape(query)
		testRequest(t, client, s.URL, query, expectedLocation)
	})
}
