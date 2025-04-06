package main

import (
	http2 "github.com/blazejsewera/sein/http"
	"github.com/blazejsewera/sein/monitor"
	"log/slog"
	"net/http"
)

func main() {
	sr := &ServiceResolver{
		defaultSearchHome: "https://sewera.cc/searxng",
		defaultSearch:     "https://sewera.cc/searxng?q={{.Query}}",
		services:          services,
	}
	sh := &http2.SearchQueryHandler{sr}
	monitor.Log().Info("starting service", slog.Int("port", 8080))
	err := http.ListenAndServe(":8080", sh)
	if err != nil {
		monitor.Log().Fatal("cannot start http server", slog.String("err", err.Error()))
	}
}
