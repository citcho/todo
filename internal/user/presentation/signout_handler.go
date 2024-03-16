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

func (sc SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Domain:  "localhost",
		Expires: time.Now(),
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}
