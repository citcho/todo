package presentation

import (
	"net/http"
	"time"

	"github.com/hexisa_go_nal_todo/internal/pkg/config"
)

type SignOutHandler struct {
}

func NewSignOutHandler() *SignOutHandler {
	return &SignOutHandler{}
}

func (soh SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()

	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   cfg.Server.ClientHost,
		Expires:  time.Now(),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
