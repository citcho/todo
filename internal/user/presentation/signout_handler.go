package presentation

import (
	"net/http"
	"time"
)

type SignOutHandler struct {
}

func NewSignOutHandler() *SignOutHandler {
	return &SignOutHandler{}
}

func (soh SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   "dev-todo.citcho.com",
		Expires:  time.Now(),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
