package command

import (
	"context"
	"fmt"

	"github.com/citcho/todo/internal/user/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u *user.User) ([]byte, error)
}

type SignInCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignIn struct {
	ur user.IUserRepository
	tg TokenGenerator
}

func NewSignIn(ur user.IUserRepository, tg TokenGenerator) *SignIn {
	return &SignIn{
		ur: ur,
		tg: tg,
	}
}

func (si *SignIn) Invoke(ctx context.Context, cmd SignInCommand) (string, error) {
	u, err := si.ur.FetchByEmail(ctx, cmd.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password()), []byte(cmd.Password)); err != nil {
		return "", err
	}

	t, err := si.tg.GenerateToken(ctx, u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	token := string(t)

	return token, nil
}
