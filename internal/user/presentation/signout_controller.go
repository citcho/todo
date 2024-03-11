package presentation

import (
	"net/http"
	"time"
)

type SignOutController struct {
}

func NewSignOutController() SignOutController {
	return SignOutController{}
}

func (sc SignOutController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
