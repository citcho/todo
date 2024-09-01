package command

import (
	"context"
	"fmt"

	"github.com/citcho/todo/internal/todo/domain/todo"
)

type CompleteCommand struct {
	Id string
}

type Complete struct {
	tr todo.ITodoRepository
}

func NewComplete(tr todo.ITodoRepository) *Complete {
	return &Complete{tr}
}

func (c *Complete) Invoke(ctx context.Context, cmd CompleteCommand) error {
	t, err := c.tr.FindById(ctx, cmd.Id)
	if err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	if err = t.Complete(ctx); err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	if err = c.tr.Update(ctx, t); err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	return nil
}
