package server

import (
	"fmt"
	"net/http"

	"github.com/citcho/todo/internal/pkg/auth"
	"github.com/citcho/todo/internal/pkg/config"
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

func corsMiddleware(cfg *config.Config) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 開発環境だと443ポートの使い回しができないので、クライアントのポートを指定する
			if cfg.AppEnv == "dev" {
				w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("https://%s:%d", cfg.ClientHost, cfg.ClientPort))
			}
			if cfg.AppEnv == "prod" {
				w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("https://%s", cfg.ClientHost))
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			h.ServeHTTP(w, r)
		})
	}
}
