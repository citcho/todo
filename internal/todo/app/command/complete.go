package command

import (
	"context"
	"fmt"

	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

type CompleteCommand struct {
	Id string
}

type CompleteHandler struct {
	tr todo.ITodoRepository
}

func NewCompleteHandler(tr todo.ITodoRepository) *CompleteHandler {
	return &CompleteHandler{
		tr: tr,
	}
}

func (sh *CompleteHandler) Handle(ctx context.Context, cmd CompleteCommand) error {
	t, err := sh.tr.FindById(ctx, cmd.Id)
	if err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	if err = t.Complete(ctx); err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	if err = sh.tr.Update(ctx, t); err != nil {
		return fmt.Errorf("Todoの完了に失敗しました。: %w", err)
	}

	return nil
}
