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

	t.Run("redirect to default search site homepage when there is no search query", func(t *testing.T) {
		query := ""
		expectedLocation := "https://example.com"
		testRequest(t, client, s.URL, query, expectedLocation)
	})

	t.Run("redirect and pass search query to default search site "+
		"when the search query doesn't include any special command or special characters", func(t *testing.T) {
		query := "abc"
		expectedLocation := "https://example.com?q=" + query
		testRequest(t, client, s.URL, query, expectedLocation)
	})

	t.Run("redirect and pass URL-escaped search query to default search site "+
		"when the search query doesn't include any special command "+
		"but includes special characters", func(t *testing.T) {
		query := "abc def @!#$/"
		expectedLocation := "https://example.com?q=" + url.QueryEscape(query)
		testRequest(t, client, s.URL, query, expectedLocation)
	})
}

func TestRedirectToExternalService(t *testing.T) {
	s := httptest.NewServer(newSearchQueryHandler())
	defer s.Close()
	client := newHTTPClient()

	t.Run("when the search query includes a bang-command (starts with an exclamation mark)", func(t *testing.T) {
		t.Run("redirect to external service's homepage "+
			"when there is no further query to the external service", func(t *testing.T) {
			query := "!w"
			expectedLocation := "https://en.wikipedia.org"
			testRequest(t, client, s.URL, query, expectedLocation)
		})

		t.Run("redirect and pass search query to external service "+
			"when there is a query to the external service", func(t *testing.T) {
			extQuery := "abc"
			query := "!w " + extQuery
			expectedLocation := "https://en.wikipedia.org/wiki/" + extQuery
			testRequest(t, client, s.URL, query, expectedLocation)
		})

		t.Run("redirect and pass URL-escaped search query to external service "+
			"when there is a query to the external service "+
			"with special characters", func(t *testing.T) {
			extQuery := "abc def @!#$/"
			query := "!w " + extQuery
			expectedLocation := "https://en.wikipedia.org/wiki/" + url.QueryEscape(extQuery)
			testRequest(t, client, s.URL, query, expectedLocation)
		})
	})

	t.Run("when the search query includes a calculation expression (starts with an equals sign)", func(t *testing.T) {
		t.Run("redirect to calculation service's homepage "+
			"when there is no further query", func(t *testing.T) {
			query := "="
			expectedLocation := "https://rinkcalc.app"
			testRequest(t, client, s.URL, query, expectedLocation)
		})

		t.Run("redirect and pass query to calculation service "+
			"when there is an expression", func(t *testing.T) {
			extQuery := "1inch to cm"
			query := "=" + extQuery
			expectedLocation := "https://rinkcalc.app?q=" + url.QueryEscape(extQuery)
			testRequest(t, client, s.URL, query, expectedLocation)
		})
	})
}
