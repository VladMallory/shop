package product_transport

import (
	"authTest/internal/features/product/domain"
	http_server "authTest/internal/server"
	"context"
	"net/http"
)

type ProductHTTPHandler struct {
	ProductService ProductService
}

type ProductService interface {
	GetProducts(ctx context.Context, limit *int, offset *int) ([]domain.Product, error)
	GetProduct(ctx context.Context, id int) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
}

func NewProductHTTPHandler(productService ProductService) *ProductHTTPHandler {
	return &ProductHTTPHandler{
		ProductService: productService,
	}
}

func (p *ProductHTTPHandler) Routes() []http_server.Route {
	return []http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: p.GetProducts,
		},
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: p.CreateProduct,
		},
	}
}
