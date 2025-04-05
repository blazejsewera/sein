package main

import (
	"fmt"
	"github.com/blazejsewera/sein/monitor"
	"log/slog"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Hello World")
}

var services = map[string]string{
	"default": "https://sewera.cc/searxng",
}

func validateServices() {
	for _, serviceUrl := range services {
		if !isParseable(serviceUrl) {
			monitor.Log().Fatal("cannot parse service URL", slog.String("serviceUrl", serviceUrl))
		}
	}
}

func isParseable(rawUrl string) bool {
	parsed, err := url.Parse(rawUrl)
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}

func routeDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services["default"], http.StatusFound)
}
