package user

import (
	"context"
	"errors"
)

type UserService struct {
	ur IUserRepository
}

func NewUserService(ur IUserRepository) *UserService {
	return &UserService{
		ur,
	}
}

func (us *UserService) Exists(ctx context.Context, u *User) (bool, error) {
	exists, err := us.ur.Exists(ctx, u)
	if err != nil {
		return exists, err
	}
	if exists {
		return exists, errors.New("既に登録されているメールアドレスです。")
	}

	return exists, nil
}
