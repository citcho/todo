package presentation

import (
	"net/http"

	"github.com/citcho/todo/internal/user/app/query"
)

type GetCurrentUserHandler struct {
	gcu *query.GetCurrentUser
}

func NewGetCurrentUserHandler(gcu *query.GetCurrentUser) *GetCurrentUserHandler {
	return &GetCurrentUserHandler{gcu}
}

func (gcuh *GetCurrentUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := gcuh.gcu.Invoke(r.Context())
	if err != nil {
		RespondJSON(r.Context(), w, err, http.StatusBadRequest)
		return
	}

	rsp := struct {
		User query.GetCurrentUserDto `json:"user"`
	}{
		User: dto,
	}

	RespondJSON(r.Context(), w, rsp, http.StatusOK)
}
