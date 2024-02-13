package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/user/app/command"
)

type SignupController struct {
	sh *command.SignupHandler
}

func NewSignupController(s *command.SignupHandler) *SignupController {
	return &SignupController{s}
}

func (s *SignupController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.SignupCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.sh.Handle(r.Context(), cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
