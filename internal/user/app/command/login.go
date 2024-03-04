package command

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u *user.User) ([]byte, error)
}

type SignInCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInHandler struct {
	ur user.IUserRepository
	tg TokenGenerator
}

func NewSignInHandler(ur user.IUserRepository, tg TokenGenerator) *SignInHandler {
	return &SignInHandler{
		ur: ur,
		tg: tg,
	}
}

func (lh *SignInHandler) Handle(ctx context.Context, cmd SignInCommand) (string, error) {
	u, err := lh.ur.FetchByEmail(ctx, cmd.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password()), []byte(cmd.Password)); err != nil {
		return "", err
	}

	t, err := lh.tg.GenerateToken(ctx, u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	token := string(t)

	return token, nil
}
