package presentation

import (
	"log"

	"github.com/citcho/todo/internal/pkg/auth"
	"github.com/citcho/todo/internal/pkg/clock"
	"github.com/citcho/todo/internal/user/adapter/mysql/repository"
	"github.com/citcho/todo/internal/user/app/command"
	"github.com/citcho/todo/internal/user/app/query"
)

type UserHandlers struct {
	GetCurrentUserHandler *GetCurrentUserHandler
	SignUpHandler         *SignUpHandler
	SignInHandler         *SignInHandler
	SignOutHandler        *SignOutHandler
}

func NewUserHandlers() *UserHandlers {
	r := repository.NewUserRepository()
	jwter, err := auth.NewJWTer(clock.RealClocker{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	ghu := query.NewGetCurrentUser(r)
	su := command.NewSignUp(r)
	si := command.NewSignIn(r, jwter)

	return &UserHandlers{
		GetCurrentUserHandler: NewGetCurrentUserHandler(ghu),
		SignUpHandler:         NewSignUpHandler(su),
		SignInHandler:         NewSignInHandler(si),
		SignOutHandler:        NewSignOutHandler(),
	}
}
