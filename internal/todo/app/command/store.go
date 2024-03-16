package command

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/pkg/auth"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

type StoreCommand struct {
	Id      string
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Store struct {
	tr todo.ITodoRepository
}

func NewStore(tr todo.ITodoRepository) *Store {
	return &Store{tr}
}

func (s *Store) Invoke(ctx context.Context, cmd StoreCommand) error {
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

	if err := s.tr.Save(ctx, u); err != nil {
		return fmt.Errorf("Todoの作成に失敗しました。: %w", err)
	}

	return nil
}
