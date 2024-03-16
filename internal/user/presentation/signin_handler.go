package presentation

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hexisa_go_nal_todo/internal/user/app/command"
)

type SignInHandler struct {
	si *command.SignIn
}

func NewSignInHandler(si *command.SignIn) *SignInHandler {
	return &SignInHandler{si}
}

func (sih *SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.SignInCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := sih.si.Invoke(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		Domain:   "dev-todo.citcho.com",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
