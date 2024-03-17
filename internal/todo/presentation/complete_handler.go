package presentation

import (
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/todo/app/command"
)

type CompleteHandler struct {
	c *command.Complete
}

func NewCompleteHandler(c *command.Complete) *CompleteHandler {
	return &CompleteHandler{c}
}

func (ch *CompleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.CompleteCommand

	cmd.Id = r.PathValue("id")

	if err := ch.c.Invoke(r.Context(), cmd); err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
