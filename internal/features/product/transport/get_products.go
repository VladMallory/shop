package product_transport

import (
	core_logger "authTest/internal/logger"
	http_response "authTest/internal/transport/http/response"
	http_utils "authTest/internal/transport/http/utils"
	"net/http"
)

type GetProductsResponse []ProductDTO

func (h *ProductHTTPHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := http_response.NewHTTPResponseHandler(logger, w)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.RespondError("failed to get query params: %w", err)
		return
	}

	products, err := h.ProductService.GetProducts(ctx, limit, offset)
	if err != nil {
		responseHandler.RespondError("failed to get products: %w", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, productsDTOFromDomains(products))
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, err
	}

	offset, err := http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, err
	}

	return limit, offset, nil
}
