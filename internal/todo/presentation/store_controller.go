package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/pkg/ulid"
	"github.com/hexisa_go_nal_todo/internal/todo/app/command"
)

type StoreController struct {
	sh *command.StoreHandler
}

func NewStoreController(s *command.StoreHandler) *StoreController {
	return &StoreController{s}
}

func (s *StoreController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.StoreCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd.Id = ulid.NewULID()

	if err := s.sh.Handle(r.Context(), cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
