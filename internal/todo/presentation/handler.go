package presentation

import (
	"github.com/hexisa_go_nal_todo/internal/todo/adapter/mysql/repository"
	"github.com/hexisa_go_nal_todo/internal/todo/app/command"
	"github.com/hexisa_go_nal_todo/internal/todo/app/query"
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
