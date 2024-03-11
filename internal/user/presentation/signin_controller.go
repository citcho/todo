package presentation

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hexisa_go_nal_todo/internal/user/app/command"
)

type SignInController struct {
	lh *command.SignInHandler
}

func NewSignInController(s *command.SignInHandler) *SignInController {
	return &SignInController{s}
}

func (lc *SignInController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var cmd command.SignInCommand

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := lc.lh.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: nginxでSSL設定後にSecureを有効にする
	cookie := http.Cookie{
		Name:    "token",
		Value:   jwt,
		Path:    "/",
		Domain:  "localhost",
		Expires: time.Now().Add(24 * time.Hour),
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
