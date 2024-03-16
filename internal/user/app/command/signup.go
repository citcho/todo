package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/pkg/encrypt"
	"github.com/hexisa_go_nal_todo/internal/pkg/ulid"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
)

type SignUpCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	ur user.IUserRepository
	us *user.UserService
}

func NewSignUp(ur user.IUserRepository) *SignUp {
	return &SignUp{
		ur: ur,
		us: user.NewUserService(ur),
	}
}

func (su *SignUp) Invoke(ctx context.Context, cmd SignUpCommand) error {
	ulid := ulid.NewULID()

	u, err := user.NewUser(
		ulid,
		cmd.Name,
		cmd.Email,
		cmd.Password,
		encrypt.Encrypter{},
	)
	if err != nil {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", err)
	}

	exists, err := su.us.Exists(ctx, u)
	if err != nil {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", err)
	}
	if exists {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", errors.New("既に登録されているメールアドレスです。"))
	}

	if err := su.ur.Save(ctx, u); err != nil {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", err)
	}

	return nil
}
