package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/citcho/todo/internal/pkg/ulid"
	"github.com/citcho/todo/internal/todo/app/command"
)

type StoreHandler struct {
	s *command.Store
}

func NewStoreHandler(s *command.Store) *StoreHandler {
	return &StoreHandler{s}
}

func (sh *StoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.StoreCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	cmd.Id = ulid.NewULID()

	if err := sh.s.Invoke(r.Context(), cmd); err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
