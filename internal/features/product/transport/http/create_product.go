package product_transport

import (
	core_logger "authTest/internal/platform/logger"
	http_request "authTest/internal/transport/http/request"
	http_response "authTest/internal/transport/http/response"
	"net/http"
)

type CreateProductRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,min=1,max=1000"`
	Price       int    `json:"price" validate:"required,min=1"`
}

type CreateProductResponse ProductDTO

func (h *ProductHTTPHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := http_response.NewHTTPResponseHandler(logger, w)

	var req CreateProductRequest

	if err := http_request.DecodeAndValidate(r, &req); err != nil {
		responseHandler.RespondError("failed to validate request", err)
		return
	}

}
