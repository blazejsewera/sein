package resolver

import (
	"bytes"
	"github.com/blazejsewera/sein/monitor"
	"log/slog"
	"net/url"
	"text/template"
)

type Resolver struct {
	defaultSearchHome  string
	defaultSearch      string
	calculationService string
	services           map[string]string
}

func New(defaultSearchHome, defaultSearch string) *Resolver {
	calculationService := "https://rinkcalc.app?q={{.Query}}"
	services := map[string]string{
		"w":   "https://en.wikipedia.org/wiki/{{.Query}}",
		"ddg": "https://duckduckgo.com?q={{.Query}}",
	}

	return &Resolver{
		defaultSearchHome,
		defaultSearch,
		calculationService,
		services,
	}
}

func (r *Resolver) ParseQueryToRedirectLocation(query string) string {
	if query == "" {
		return r.defaultSearchHome
	}

	return renderTemplateURI(r.defaultSearch, query)
}

func renderTemplateURI(templateString string, query string) string {
	templateName := "t"
	t, err := template.New(templateName).Parse(templateString)
	if err != nil {
		monitor.Log().Error("cannot parse template", slog.String("template", templateString), slog.String("query", query), slog.String("err", err.Error()))
		return ""
	}

	b := bytes.Buffer{}
	err = t.ExecuteTemplate(&b, templateName, wrapQueryForTemplate(query))
	if err != nil {
		monitor.Log().Error("cannot execute template with query", slog.String("template", templateString), slog.String("query", query), slog.String("err", err.Error()))
		return ""
	}
	return b.String()
}

func wrapQueryForTemplate(query string) any {
	return struct{ Query string }{url.QueryEscape(query)}
}
