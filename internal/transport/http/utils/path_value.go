package http_utils

import (
	"authTest/internal/errs"
	"fmt"
	"net/http"
	"strconv"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	val := r.PathValue(key)
	if val == "" {
		return 0, nil
	}

	rVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("invalid path value: %s :%w", val, errs.ErrInvalidArgument)
	}

	return rVal, nil
}
