package user

import (
	"context"
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
		return false, err
	}

	return exists, nil
}
