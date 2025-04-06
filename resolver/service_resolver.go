package resolver

type Resolver struct {
	DefaultSearchHome string
	DefaultSearch     string
	Services          map[string]string
}

func New(defaultSearchHome, defaultSearch string) *Resolver {
	services := map[string]string{
		"w":   "https://en.wikipedia.org/wiki/{{.Query}}",
		"ddg": "https://duckduckgo.com?q={{.Query}}",
	}

	return &Resolver{
		defaultSearchHome,
		defaultSearch,
		services,
	}
}
