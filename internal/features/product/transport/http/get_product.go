package product_transport

import (
	core_logger "authTest/internal/platform/logger"
	http_response "authTest/internal/transport/http/response"
	http_utils "authTest/internal/transport/http/utils"
	"net/http"
)

type GetProductResponse ProductDTO

func (h *ProductHTTPHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := http_response.NewHTTPResponseHandler(logger, w)

	id, err := http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get path value", err)
		return
	}

	product, err := h.ProductService.GetProduct(ctx, id)
	if err != nil {
		responseHandler.RespondError("failed to get product", err)
		return
	}

	response := GetProductResponse(productDTOFromDomain(product))
	responseHandler.RespondJSON(http.StatusOK, response)
}
