package presentation

import (
	"net/http"

	"github.com/hexisa_go_nal_todo/internal/user/app/query"
)

type GetCurrentUserController struct {
	gcuh *query.GetCurrentUserHandler
}

func NewGetCurrentUserController(gcuh *query.GetCurrentUserHandler) *GetCurrentUserController {
	return &GetCurrentUserController{gcuh}
}

func (gcuc *GetCurrentUserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := gcuc.gcuh.Handle(r.Context())
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
