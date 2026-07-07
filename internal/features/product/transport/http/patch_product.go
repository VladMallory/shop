package product_transport

import (
	"authTest/internal/errs"
	"authTest/internal/features/product/domain"
	core_logger "authTest/internal/platform/logger"
	http_request "authTest/internal/transport/http/request"
	http_response "authTest/internal/transport/http/response"
	http_types "authTest/internal/transport/http/types"
	http_utils "authTest/internal/transport/http/utils"
	"fmt"
	"net/http"
	"unicode/utf8"
)

type PatchProductRequest struct {
	Title       http_types.Nullable[string] `json:"title"`
	Description http_types.Nullable[string] `json:"description"`
	Price       http_types.Nullable[int]    `json:"price"`
}

type PatchProductResponse ProductDTO

func (r *PatchProductRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title cannot be set to NULL: %w", errs.ErrInvalidArgument)
		}

		titleLength := utf8.RuneCountInString(*r.Title.Value)
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf("title must be between 1 and 100: %w", errs.ErrInvalidArgument)
		}
	}

	if r.Description.Set {
		if r.Description.Value == nil {
			return fmt.Errorf("description cannot be set to NULL: %w", errs.ErrInvalidArgument)
		}

		descriptionLength := utf8.RuneCountInString(*r.Description.Value)
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("description must be between 1 and 100: %w", errs.ErrInvalidArgument)
		}
	}

	if r.Price.Set {
		if r.Price.Value == nil {
			return fmt.Errorf("price cannot be set to NULL: %w", errs.ErrInvalidArgument)
		}

		if *r.Price.Value <= 0 {
			return fmt.Errorf("price must be greater than 0: %w", errs.ErrInvalidArgument)
		}
	}

	return nil
}

func (h *ProductHTTPHandler) PathProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := http_response.NewHTTPResponseHandler(logger, w)

	id, err := http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get id path value: %w", err)
		return
	}

	var productPatchRequest PatchProductRequest
	if err := http_request.DecodeAndValidate(r, &productPatchRequest); err != nil {
		responseHandler.RespondError("failed to decode reques: %w", err)
		return
	}

	if err := productPatchRequest.Validate(); err != nil {
		responseHandler.RespondError("validation error: %w", err)
		return
	}

}

func productPatchFromRequest(r PatchProductRequest) domain.ProductPatch {
	return domain.NewProductPatch(
		r.Title.ToDomain(),
		r.Description.ToDomain(),
		r.Price.ToDomain(),
	)
}
