package search

import (
	"bytes"
	"github.com/blazejsewera/sein/monitor"
	"log/slog"
	"net/url"
	"text/template"
)

type Service struct {
	homepage          string
	searchURITemplate string
}

func NewService(homepage, searchURITemplate string) Service {
	return Service{homepage, searchURITemplate}
}

func (s Service) Homepage() string {
	return s.homepage
}

func (s Service) RenderTemplateURI(searchQuery string) string {
	templateName := "t"
	t, err := template.New(templateName).Parse(s.searchURITemplate)
	if err != nil {
		monitor.Log().Error("cannot parse template",
			slog.String("template", s.searchURITemplate),
			slog.String("searchQuery", searchQuery),
			slog.String("err", err.Error()))
		return ""
	}

	b := bytes.Buffer{}
	err = t.ExecuteTemplate(&b, templateName, wrapQueryForTemplate(searchQuery))
	if err != nil {
		monitor.Log().Error("cannot execute template with query",
			slog.String("template", s.searchURITemplate),
			slog.String("searchQuery", searchQuery),
			slog.String("err", err.Error()))
		return ""
	}
	return b.String()
}

func wrapQueryForTemplate(query string) any {
	return struct{ Query string }{url.QueryEscape(query)}
}
