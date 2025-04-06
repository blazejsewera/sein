package http

import (
	"bytes"
	"github.com/blazejsewera/sein/resolver"
	"log/slog"
	"net/http"
	"net/url"
	"text/template"

	"github.com/blazejsewera/sein/monitor"
)

type SearchQueryHandler struct {
	resolver *resolver.Resolver
}

func NewSearchQueryHandler(resolver *resolver.Resolver) *SearchQueryHandler {
	return &SearchQueryHandler{resolver}
}

var _ http.Handler = new(SearchQueryHandler)

func (s *SearchQueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		msg := "cannot parse request URI"
		monitor.Log().Error(msg, slog.String("uri", r.RequestURI))
		http.Error(w, msg+": "+r.RequestURI, http.StatusBadRequest)
		return
	}
	query := uri.Query()
	if !query.Has("q") {
		http.Redirect(w, r, s.resolver.DefaultSearchHome, http.StatusSeeOther)
		return
	}
	q := query.Get("q")
	http.Redirect(w, r, renderTemplateURI(s.resolver.DefaultSearch, q), http.StatusSeeOther)
}

func renderTemplateURI(templateString string, q string) string {
	templateName := "t"
	t, err := template.New(templateName).Parse(templateString)
	if err != nil {
		monitor.Log().Error("cannot parse template", slog.String("template", templateString), slog.String("query", q), slog.String("err", err.Error()))
		return ""
	}
	b := bytes.Buffer{}
	err = t.ExecuteTemplate(&b, templateName, struct{ Query string }{url.QueryEscape(q)})
	if err != nil {
		monitor.Log().Error("cannot execute template with query", slog.String("template", templateString), slog.String("query", q), slog.String("err", err.Error()))
		return ""
	}
	return b.String()
}
