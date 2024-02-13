package user

import (
	"context"
)

//go:generate mockgen -source=./repository.go -destination=./mock/user_repository.go -package=mock
type IUserRepository interface {
	Save(context.Context, *User) error
	Exists(context.Context, *User) (bool, error)
	FetchByEmail(context.Context, string) (*User, error)
	FetchByUlid(context.Context, string) (*User, error)
}
