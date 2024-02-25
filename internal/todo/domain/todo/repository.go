package todo

import (
	"context"
)

//go:generate mockgen -source=./repository.go -destination=./mock/todo_repository.go -package=mock
type ITodoRepository interface {
	Save(context.Context, *Todo) error
	FindById(context.Context, string) (*Todo, error)
	Update(context.Context, *Todo) error
	FindAll(context.Context, string) ([]*Todo, error)
}
