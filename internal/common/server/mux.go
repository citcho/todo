package server

import (
	"context"
	"log"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
	"github.com/hexisa_go_nal_todo/internal/common/clock"
	"github.com/hexisa_go_nal_todo/internal/common/config"
	"github.com/hexisa_go_nal_todo/internal/common/database"
	todo_repository "github.com/hexisa_go_nal_todo/internal/todo/adapter/mysql/repository"
	todo_command "github.com/hexisa_go_nal_todo/internal/todo/app/command"
	todo_query "github.com/hexisa_go_nal_todo/internal/todo/app/query"
	todo_presentation "github.com/hexisa_go_nal_todo/internal/todo/presentation"
	user_repository "github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/repository"
	user_command "github.com/hexisa_go_nal_todo/internal/user/app/command"
	user_query "github.com/hexisa_go_nal_todo/internal/user/app/query"
	user_presentation "github.com/hexisa_go_nal_todo/internal/user/presentation"
)

func NewMux(ctx context.Context, cfg *config.Config) (*http.ServeMux, func(), error) {
	db, cleanup, err := database.NewDB(ctx, cfg.DB)
	if err != nil {
		return nil, cleanup, err
	}

	mux := http.NewServeMux()

	jwter, err := auth.NewJWTer(clock.RealClocker{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	userRepository := user_repository.NewUserRepository(db)

	signupHandler := user_command.NewSignupHandler(userRepository)
	mux.Handle("POST /signup", user_presentation.NewSignupController(signupHandler))

	loginHandler := user_command.NewLoginHandler(userRepository, jwter)
	mux.Handle("POST /login", user_presentation.NewLoginController(loginHandler))

	getCurrentUserHandler := user_query.NewGetCurrentUserHandler(userRepository)
	getCurrentUserController := user_presentation.NewGetCurrentUserController(getCurrentUserHandler)

	mux.Handle("/me", with(getCurrentUserController, jwtMiddleware(jwter)))

	todoRepository := todo_repository.NewTodoRepository(db)

	storeHandler := todo_command.NewStoreHandler(todoRepository)
	storeController := todo_presentation.NewStoreController(storeHandler)
	mux.Handle("POST /todos", with(storeController, jwtMiddleware(jwter)))

	completeHandler := todo_command.NewCompleteHandler(todoRepository)
	completeController := todo_presentation.NewCompleteController(completeHandler)
	mux.Handle("PATCH /todos/{id}/complete", with(completeController, jwtMiddleware(jwter)))

	getTodosHandler := todo_query.NewGetTodosHandler(todoRepository)
	getTodosController := todo_presentation.NewGetTodosController(getTodosHandler)
	mux.Handle("GET /todos", with(getTodosController, jwtMiddleware(jwter)))

	return mux, cleanup, nil
}
