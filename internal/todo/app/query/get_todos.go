package query

import (
	"context"
	"fmt"

	"github.com/citcho/todo/internal/pkg/auth"
	"github.com/citcho/todo/internal/todo/domain/todo"
)

type GetTodosDto struct {
	Todos []*todo.Todo
}

type GetTodos struct {
	tr todo.ITodoRepository
}

func NewGetTodos(tr todo.ITodoRepository) *GetTodos {
	return &GetTodos{tr}
}

func (gt *GetTodos) Invoke(ctx context.Context) ([]*todo.Todo, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("Todoの取得に失敗しました。: %s", "ユーザーIDが取得できませんでした。")
	}

	dto, err := gt.tr.FindAll(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("Todoの取得に失敗しました。: %w", err)
	}

	return dto, nil
}
