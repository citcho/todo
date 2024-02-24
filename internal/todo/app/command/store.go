package command

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

type StoreCommand struct {
	Id      string
	Title   string `json:"title"`
	Content string `json:"content"`
}

type StoreHandler struct {
	tr todo.ITodoRepository
}

func NewStoreHandler(tr todo.ITodoRepository) *StoreHandler {
	return &StoreHandler{
		tr: tr,
	}
}

func (sh *StoreHandler) Handle(ctx context.Context, cmd StoreCommand) error {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("Todoの作成に失敗しました。: %s", "ユーザーIDが取得できませんでした。")
	}

	u, err := todo.NewTodo(
		cmd.Id,
		userId,
		cmd.Title,
		cmd.Content,
	)
	if err != nil {
		return fmt.Errorf("Todoの作成に失敗しました。: %w", err)
	}

	if err := sh.tr.Save(ctx, u); err != nil {
		return fmt.Errorf("Todoの作成に失敗しました。: %w", err)
	}

	return nil
}
