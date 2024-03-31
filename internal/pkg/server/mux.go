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

type preflightHandler struct{}

func (ph preflightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewMux(ctx context.Context, cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	jwter, err := auth.NewJWTer(clock.RealClocker{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	userHandlers := user_presentation.NewUserHandlers()
	todoHandlers := todo_presentation.NewTodoHandlers()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("OPTIONS /{path...}", with(preflightHandler{}, corsMiddleware(cfg)))
	mux.Handle("POST /signup", with(userHandlers.SignUpHandler, corsMiddleware(cfg)))
	mux.Handle("POST /signin", with(userHandlers.SignInHandler, corsMiddleware(cfg)))
	mux.Handle("POST /signout", with(userHandlers.SignOutHandler, jwtMiddleware(jwter), corsMiddleware(cfg)))
	mux.Handle("GET /me", with(userHandlers.GetCurrentUserHandler, jwtMiddleware(jwter), corsMiddleware(cfg)))
	mux.Handle("POST /todos", with(todoHandlers.StoreHandler, jwtMiddleware(jwter), corsMiddleware(cfg)))
	mux.Handle("PATCH /todos/{id}/complete", with(todoHandlers.CompleteHandler, jwtMiddleware(jwter), corsMiddleware(cfg)))
	mux.Handle("GET /todos", with(todoHandlers.GetTodosHandler, jwtMiddleware(jwter), corsMiddleware(cfg)))

	return mux
}
