package server

import (
	"context"
	"log"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/pkg/auth"
	"github.com/hexisa_go_nal_todo/internal/pkg/clock"
	"github.com/hexisa_go_nal_todo/internal/pkg/config"
	todo_presentation "github.com/hexisa_go_nal_todo/internal/todo/presentation"
	user_presentation "github.com/hexisa_go_nal_todo/internal/user/presentation"
)

func NewMux(ctx context.Context, cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	jwter, err := auth.NewJWTer(clock.RealClocker{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	userHandlers := user_presentation.NewUserHandlers()
	todoHandlers := todo_presentation.NewTodoHandlers()

	mux.HandleFunc("OPTIONS /{path...}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cfg.Server.ClientUrl)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("POST /signup", with(userHandlers.SignUpHandler, corsMiddleware(cfg.Server)))
	mux.Handle("POST /signin", with(userHandlers.SignInHandler, corsMiddleware(cfg.Server)))
	mux.Handle("POST /signout", with(userHandlers.SignOutHandler, jwtMiddleware(jwter), corsMiddleware(cfg.Server)))
	mux.Handle("GET /me", with(userHandlers.GetCurrentUserHandler, jwtMiddleware(jwter), corsMiddleware(cfg.Server)))
	mux.Handle("POST /todos", with(todoHandlers.StoreHandler, jwtMiddleware(jwter), corsMiddleware(cfg.Server)))
	mux.Handle("PATCH /todos/{id}/complete", with(todoHandlers.CompleteHandler, jwtMiddleware(jwter), corsMiddleware(cfg.Server)))
	mux.Handle("GET /todos", with(todoHandlers.GetTodosHandler, jwtMiddleware(jwter), corsMiddleware(cfg.Server)))

	return mux
}
