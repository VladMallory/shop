package http_response

import (
	"authTest/internal/errs"
	logger "authTest/internal/platform/logger"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type HTTPResponseHandler struct {
	logger *logger.Logger
	w      http.ResponseWriter
}

func NewHTTPResponseHandler(
	logger *logger.Logger,
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		logger: logger,
		w:      w,
	}
}

func (h *HTTPResponseHandler) RespondError(msg string, err error) {
	var (
		statusCode int
		logFunc    func(string, ...any)
	)

	if h.logger == nil {
		return
	}

	switch {
	case errors.Is(err, errs.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.logger.Debug
	case errors.Is(err, errs.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.logger.Warn
	case errors.Is(err, errs.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.logger.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.logger.Error
	}

	logFunc(msg, slog.Any("error", err))
	h.respondError(statusCode, msg, err)
}

func (h *HTTPResponseHandler) RespondJSON(statusCode int, data any) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(data); err != nil {
		h.logger.Error("failed to encode response data", slog.Any("error", err))
	}
}

func (h *HTTPResponseHandler) respondError(statusCode int, msg string, err error) {
	h.w.WriteHeader(statusCode)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	message := map[string]string{
		"message": msg,
		"error":   errMsg,
	}

	if err := json.NewEncoder(h.w).Encode(message); err != nil {
		h.logger.Error("failed to encode response data", slog.Any("error", err))
	}
}

func (h *HTTPResponseHandler) RespondNoContent() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) RespondPanic(p any, msg string) {
	err := fmt.Errorf("unexpected panic: %v", p)
	h.logger.Error(msg, slog.Any("error", err))

	h.respondError(http.StatusInternalServerError, msg, err)
}
