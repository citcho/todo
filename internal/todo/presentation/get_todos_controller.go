package presentation

import (
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/todo/app/query"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

type GetTodosController struct {
	gth *query.GetTodosHandler
}

type GetTodosResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"isComplete"`
}

func NewGetTodosController(gth *query.GetTodosHandler) *GetTodosController {
	return &GetTodosController{gth}
}

func (gtc *GetTodosController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := gtc.gth.Handle(r.Context())
	if err != nil {
		RespondJSON(r.Context(), w, err, http.StatusBadRequest)
		return
	}

	rsp := struct {
		Todos []GetTodosResponse `json:"todos"`
	}{
		Todos: toResponse(todos),
	}

	RespondJSON(r.Context(), w, rsp, http.StatusOK)
}

func toResponse(todos []*todo.Todo) []GetTodosResponse {
	var res []GetTodosResponse
	for _, t := range todos {
		res = append(res, GetTodosResponse{
			ID:         t.Id(),
			Title:      t.Title(),
			Content:    t.Content(),
			IsComplete: t.IsComplete(),
		})
	}
	return res
}
