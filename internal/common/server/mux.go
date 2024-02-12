package server

import (
	"context"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/common/config"
	"github.com/hexisa_go_nal_todo/internal/common/database"
	"github.com/hexisa_go_nal_todo/internal/user/adapter/mysql/repository"
	"github.com/hexisa_go_nal_todo/internal/user/app/command"
	"github.com/hexisa_go_nal_todo/internal/user/presentation"
)

func NewMux(ctx context.Context, cfg *config.Config) (*http.ServeMux, func(), error) {
	db, cleanup, err := database.NewDB(ctx, cfg.DB)
	if err != nil {
		return nil, cleanup, err
	}

	mux := http.NewServeMux()

	ur := repository.NewUserRepository(db)
	sh := command.NewSignupHandler(ur)
	mux.Handle("POST /users", presentation.NewSignupController(sh))

	return mux, cleanup, nil
}
