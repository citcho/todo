package presentation

import (
	"github.com/citcho/todo/internal/todo/adapter/mysql/repository"
	"github.com/citcho/todo/internal/todo/app/command"
	"github.com/citcho/todo/internal/todo/app/query"
)

type TodoHandlers struct {
	GetTodosHandler *GetTodosHandler
	CompleteHandler *CompleteHandler
	StoreHandler    *StoreHandler
}

func NewTodoHandlers() *TodoHandlers {
	r := repository.NewTodoRepository()

	gt := query.NewGetTodos(r)
	c := command.NewComplete(r)
	s := command.NewStore(r)

	return &TodoHandlers{
		GetTodosHandler: NewGetTodosHandler(gt),
		CompleteHandler: NewCompleteHandler(c),
		StoreHandler:    NewStoreHandler(s),
	}
}
