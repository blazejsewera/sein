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
		cmd, query, found := strings.Cut(searchQuery, " ")
		serviceKey := strings.TrimPrefix(cmd, "!")
		service, exists := r.services[serviceKey]
		if !exists {
			return r.defaultSearchService.RenderTemplateURI(searchQuery)
		}

		if !found || query == "" {
			return service.Homepage()
		}
		return service.RenderTemplateURI(query)
	}

	return r.defaultSearchService.RenderTemplateURI(searchQuery)
}
