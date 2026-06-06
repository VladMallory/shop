package http_request

import (
	"authTest/internal/errs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidate(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}

	v, ok := dest.(validatable)
	var err error
	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("request validation error: %v %w", err, errs.ErrInvalidArgument)
	}

	return nil
}
