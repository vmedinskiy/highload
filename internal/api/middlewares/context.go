package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UrlHijack(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		r.URL.Path = rctx.RoutePath
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
