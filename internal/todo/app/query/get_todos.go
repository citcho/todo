package query

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

type GetTodosDto struct {
	Todos []*todo.Todo `json:"todos"`
}

type GetTodosHandler struct {
	tr todo.ITodoRepository
}

func NewGetTodosHandler(tr todo.ITodoRepository) *GetTodosHandler {
	return &GetTodosHandler{
		tr: tr,
	}
}

func (gth *GetTodosHandler) Handle(ctx context.Context) ([]*todo.Todo, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("Todoの取得に失敗しました。: %s", "ユーザーIDが取得できませんでした。")
	}

	dto, err := gth.tr.FindAll(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("Todoの取得に失敗しました。: %w", err)
	}

	return dto, nil
}
