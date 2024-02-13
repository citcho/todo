package server

import (
	"context"
	"log"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
	"github.com/hexisa_go_nal_todo/internal/common/clock"
	"github.com/hexisa_go_nal_todo/internal/common/config"
	"github.com/hexisa_go_nal_todo/internal/common/database"
	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/repository"
	"github.com/hexisa_go_nal_todo/internal/user/app/command"
	"github.com/hexisa_go_nal_todo/internal/user/app/query"
	"github.com/hexisa_go_nal_todo/internal/user/presentation"
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

	ur := repository.NewUserRepository(db)

	sh := command.NewSignupHandler(ur)
	mux.Handle("POST /signup", presentation.NewSignupController(sh))

	lh := command.NewLoginHandler(ur, jwter)
	mux.Handle("POST /login", presentation.NewLoginController(lh))

	gcuh := query.NewGetCurrentUserHandler(ur)
	gcuc := presentation.NewGetCurrentUserController(gcuh)

	mux.Handle("/me", with(gcuc, jwtMiddleware(jwter)))

	return mux, cleanup, nil
}
