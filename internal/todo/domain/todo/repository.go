package todo

import (
	"context"
)

//go:generate mockgen -source=./repository.go -destination=./mock/todo_repository.go -package=mock
type ITodoRepository interface {
	Save(context.Context, *Todo) error
}
