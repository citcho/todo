package presentation

import (
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/todo/app/command"
)

type CompleteController struct {
	sh *command.CompleteHandler
}

func NewCompleteController(s *command.CompleteHandler) *CompleteController {
	return &CompleteController{s}
}

func (s *CompleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.CompleteCommand

	cmd.Id = r.PathValue("id")

	if err := s.sh.Handle(r.Context(), cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
