package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/user/app/command"
)

type SignUpHandler struct {
	su *command.SignUp
}

func NewSignUpHandler(su *command.SignUp) *SignUpHandler {
	return &SignUpHandler{su}
}

func (suh *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.SignUpCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := suh.su.Invoke(r.Context(), cmd); err != nil {
		rsp := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		RespondJSON(r.Context(), w, rsp, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
