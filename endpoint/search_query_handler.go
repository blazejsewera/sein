package endpoint

import (
	"github.com/blazejsewera/sein/monitor"
	"github.com/blazejsewera/sein/resolver"
	"log/slog"
	"net/http"
	"net/url"
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
	queries := uri.Query()
	q := queries.Get("q")
	http.Redirect(w, r, s.resolver.ParseQueryToRedirectLocation(q), http.StatusSeeOther)
}
