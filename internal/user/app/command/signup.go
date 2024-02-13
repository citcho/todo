package command

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/pkg/encrypt"
	"github.com/hexisa_go_nal_todo/internal/pkg/ulid"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
)

type SignupCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupHandler struct {
	ur user.IUserRepository
	us *user.UserService
}

func NewSignupHandler(ur user.IUserRepository) *SignupHandler {
	return &SignupHandler{
		ur: ur,
		us: user.NewUserService(ur),
	}
}

func (sh *SignupHandler) Handle(ctx context.Context, cmd SignupCommand) error {
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

	exists, err := sh.us.Exists(ctx, u)
	if exists || err != nil {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", err)
	}

	if err := sh.ur.Save(ctx, u); err != nil {
		return fmt.Errorf("ユーザーの作成に失敗しました。: %w", err)
	}

	return nil
}
