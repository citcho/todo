package user

import (
	"context"
)

//go:generate mockgen -source=./service.go -destination=./mock/user_repository.go -package=mock
type IUserRepository interface {
	Exists(context.Context, *User) (bool, error)
}

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

	return exists, nil
}
