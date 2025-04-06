package resolver

import (
	"github.com/blazejsewera/sein/resolver/search"
	"strings"
)

type Resolver struct {
	defaultSearchService search.Service
	calculationService   search.Service
	services             map[string]search.Service
}

func New(defaultSearchHome, defaultSearch string) *Resolver {
	calculationService := search.NewService("https://rinkcalc.app", "https://rinkcalc.app?q={{.Query}}")
	services := map[string]search.Service{
		"w":   search.NewService("https://en.wikipedia.org", "https://en.wikipedia.org/wiki/{{.Query}}"),
		"ddg": search.NewService("https://duckduckgo.com", "https://duckduckgo.com?q={{.Query}}"),
	}

	return &Resolver{
		defaultSearchService: search.NewService(defaultSearchHome, defaultSearch),
		calculationService:   calculationService,
		services:             services,
	}
}

func (r *Resolver) ParseSearchQueryToRedirectLocation(searchQuery string) string {
	if searchQuery == "" {
		return r.defaultSearchService.Homepage()
	}

	if strings.HasPrefix(searchQuery, "!") {
		extQuery := strings.TrimPrefix(searchQuery, "!")
		return r.parseBangCommand(extQuery)
	} else if strings.HasPrefix(searchQuery, "=") {
		expression := strings.TrimSpace(strings.TrimPrefix(searchQuery, "="))
		return r.parseCalculationExpression(expression)
	}

	return r.defaultSearchService.RenderTemplateURI(searchQuery)
}

func (r *Resolver) parseBangCommand(searchQuery string) string {
	cmd, query, found := strings.Cut(searchQuery, " ")
	service, exists := r.services[cmd]
	if !exists {
		return r.defaultSearchService.RenderTemplateURI(searchQuery)
	}
	if !found || query == "" {
		return service.Homepage()
	}
	return service.RenderTemplateURI(query)
}

func (r *Resolver) parseCalculationExpression(expression string) string {
	if expression == "" {
		return r.calculationService.Homepage()
	}
	return r.calculationService.RenderTemplateURI(expression)
}
