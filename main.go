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
		defaultSearch:     template.New("https://sewera.cc/searxng?q={{.Query}}"),
		services:          services(),
	}
	sh := &serviceHandler{sr}
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
	}
	q := query.Get("q")
	http.Redirect(w, r, renderTemplateURI(s.resolver.defaultSearch, q), http.StatusSeeOther)
}

type Query struct {
	Query string
}

func renderTemplateURI(t *template.Template, q string) string {
	b := bytes.Buffer{}
	err := t.Execute(&b, Query{q})
	if err != nil {
		monitor.Log().Error("cannot execute template with query", slog.String("query", q))
		return ""
	}
	return b.String()
}

var _ http.Handler = new(serviceHandler)

type serviceResolver struct {
	defaultSearchHome string
	defaultSearch     *template.Template
	services          map[string]*template.Template
}

func services() map[string]*template.Template {
	wikipedia := template.New("https://en.wikipedia.org/wiki/{{.Query}}")
	duckduckgo := template.New("https://duckduckgo.com?q={{.Query}}")

	return map[string]*template.Template{
		"w":   wikipedia,
		"ddg": duckduckgo,
	}
}
