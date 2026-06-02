package handler

import (
	"authTest/internal/domain"
	"authTest/internal/errs"
	"authTest/internal/service"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "неверный формат запроса")

		return
	}

	resp, err := h.authService.Register(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrUserExists):
			respondWithError(w, http.StatusConflict, errs.ErrUserExists.Error())

		default:
			respondWithError(w, http.StatusInternalServerError, errs.ErrRegistration.Error())
		}

		return
	}

	respondWithJSON(w, http.StatusCreated, resp)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "неверный формат запроса")

		return
	}

	resp, err := h.authService.Login(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrInvalidCredentials):
			respondWithError(w, http.StatusUnauthorized, errs.ErrInvalidCredentials.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, errs.ErrRegistration.Error())
		}

		return
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func respondWithJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	respondWithJSON(w, status, map[string]string{"error": msg})
}
