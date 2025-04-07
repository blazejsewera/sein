package main

import (
	"github.com/blazejsewera/sein/endpoint"
	"github.com/blazejsewera/sein/monitor"
	"github.com/blazejsewera/sein/resolver"
	"log/slog"
	gohttp "net/http"
)

func main() {
	r := resolver.New("https://sewera.cc/searxng", "https://sewera.cc/searxng?q={{.Query}}")
	searchQueryHandler := endpoint.NewSearchQueryHandler(r)
	monitor.Log().Info("starting service", slog.Int("port", 8080))
	err := gohttp.ListenAndServe("0.0.0.0:8080", searchQueryHandler)
	if err != nil {
		monitor.Log().Fatal("cannot start http server", slog.String("err", err.Error()))
	}
}
