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

type LoginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHandler struct {
	ur user.IUserRepository
	tg TokenGenerator
}

func NewLoginHandler(ur user.IUserRepository, tg TokenGenerator) *LoginHandler {
	return &LoginHandler{
		ur: ur,
		tg: tg,
	}
}

func (lh *LoginHandler) Handle(ctx context.Context, cmd LoginCommand) (string, error) {
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
