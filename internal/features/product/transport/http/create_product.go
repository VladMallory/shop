package product_transport

import (
	"authTest/internal/features/product/domain"
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

	userDomain := domainFromRequest(req)
	userDomain, err := h.ProductService.CreateProduct(ctx, userDomain)
	if err != nil {
		responseHandler.RespondError("failed to create product", err)
		return
	}

	response := productDTOFromDomain(userDomain)
	responseHandler.RespondJSON(http.StatusCreated, CreateProductResponse(response))
}

func domainFromRequest(request CreateProductRequest) domain.Product {
	return domain.NewUninitializedProduct(request.Title, request.Description, request.Price)
}
