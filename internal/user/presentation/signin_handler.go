package presentation

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/citcho/todo/internal/pkg/config"
	"github.com/citcho/todo/internal/user/app/command"
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
		RespondJSON(r.Context(), w, ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	jwt, err := sih.si.Invoke(r.Context(), cmd)
	if err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	cfg := config.NewConfig()

	cookie := http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		Domain:   cfg.Server.ClientHost,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
