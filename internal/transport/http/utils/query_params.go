package http_utils

import (
	"authTest/internal/errs"
	"fmt"
	"net/http"
	"strconv"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return nil, nil
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return nil, fmt.Errorf("query param %s is not int: %w", key, errs.ErrInvalidArgument)
	}

	return &intVal, nil
}
