package main

import (
	"bytes"
	"github.com/blazejsewera/sein/monitor"
	"log/slog"
	"net/http"
	"net/url"
	"text/template"
)

func main() {
	sr := &serviceResolver{
		defaultSearchHome: "https://sewera.cc/searxng",
		defaultSearch:     "https://sewera.cc/searxng?q={{.Query}}",
		services:          services(),
	}
	sh := &serviceHandler{sr}
	monitor.Log().Info("starting service", slog.Int("port", 8080))
	err := http.ListenAndServe(":8080", sh)
	if err != nil {
		monitor.Log().Fatal("cannot start http server", slog.String("err", err.Error()))
	}
}

type serviceHandler struct {
	resolver *serviceResolver
}

func (s *serviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		msg := "cannot parse request URI"
		monitor.Log().Error(msg, slog.String("uri", r.RequestURI))
		http.Error(w, msg+": "+r.RequestURI, http.StatusBadRequest)
		return
	}
	query := uri.Query()
	if !query.Has("q") {
		http.Redirect(w, r, s.resolver.defaultSearchHome, http.StatusSeeOther)
		return
	}
	q := query.Get("q")
	http.Redirect(w, r, renderTemplateURI(s.resolver.defaultSearch, q), http.StatusSeeOther)
}

type Query struct {
	Query string
}

func renderTemplateURI(templateString string, q string) string {
	templateName := "t"
	t, err := template.New(templateName).Parse(templateString)
	if err != nil {
		monitor.Log().Error("cannot parse template", slog.String("template", templateString), slog.String("query", q), slog.String("err", err.Error()))
		return ""
	}
	b := bytes.Buffer{}
	err = t.ExecuteTemplate(&b, templateName, Query{q})
	if err != nil {
		monitor.Log().Error("cannot execute template with query", slog.String("template", templateString), slog.String("query", q), slog.String("err", err.Error()))
		return ""
	}
	return b.String()
}

var _ http.Handler = new(serviceHandler)

type serviceResolver struct {
	defaultSearchHome string
	defaultSearch     string
	services          map[string]string
}

func services() map[string]string {
	wikipedia := "https://en.wikipedia.org/wiki/{{.Query}}"
	duckduckgo := "https://duckduckgo.com?q={{.Query}}"

	return map[string]string{
		"w":   wikipedia,
		"ddg": duckduckgo,
	}
}
