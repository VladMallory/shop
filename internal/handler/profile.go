package handler

import (
	"authTest/internal/middleware"
	"authTest/internal/service"
	"log/slog"
	"net/http"
)

type ProfileHandler struct {
	userService *service.UserService
}

func NewProfileHandler(userService *service.UserService) *ProfileHandler {
	return &ProfileHandler{
		userService: userService,
	}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.UserIDFromCtx(r.Context())

	user, err := h.userService.GetByID(r.Context(), userID)
	if err != nil {
		slog.Error("get profile", "error", err)
		respondWithError(w, http.StatusNotFound, "пользователь не найден")

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
