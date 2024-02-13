package server

import (
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
)

type Middleware interface {
	ServeNext(h http.Handler) http.Handler
}

type MiddlewareFunc func(h http.Handler) http.Handler

func (f MiddlewareFunc) ServeNext(h http.Handler) http.Handler {
	return f(h)
}

func with(h http.Handler, ms ...Middleware) http.Handler {
	for _, m := range ms {
		h = m.ServeNext(h)
	}
	return h
}

func jwtMiddleware(j *auth.JWTer) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r, err := j.FillContext(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
