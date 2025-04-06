package sein_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func testRequest(t testing.TB, client *http.Client, serverURI, query, expectedLocation string) {
	t.Helper()

	var uri string
	if query != "" {
		q := url.QueryEscape(query)
		uri = fmt.Sprintf("%s?q=%s", serverURI, q)
	} else {
		uri = serverURI
	}
	resp, err := client.Get(uri)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Header.Get("Location") != expectedLocation {
		t.Fatalf("'Location' header mismatch: actual: %s; expected: %s", resp.Header.Get("Location"), expectedLocation)
	}
}
