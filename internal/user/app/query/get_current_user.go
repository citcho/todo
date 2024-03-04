package query

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
)

type GetCurrentUserHandler struct {
	ur user.IUserRepository
}

type GetCurrentUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGetCurrentUserHandler(ur user.IUserRepository) *GetCurrentUserHandler {
	return &GetCurrentUserHandler{
		ur: ur,
	}
}

func (gcuh *GetCurrentUserHandler) Handle(ctx context.Context) (GetCurrentUserDto, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return GetCurrentUserDto{}, fmt.Errorf("ユーザーが存在しません")
	}

	u, err := gcuh.ur.FetchById(ctx, userId)
	if err != nil {
		return GetCurrentUserDto{}, fmt.Errorf("対象のユーザーが存在しません")
	}

	dto := GetCurrentUserDto{
		Name:  u.Name(),
		Email: u.Email(),
	}

	return dto, nil
}
