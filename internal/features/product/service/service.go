package product_service

import (
	"authTest/internal/domain"
	"context"
)

type ProductService struct {
	productRepository ProductRepository
}

type ProductRepository interface {
	GetProducts(ctx context.Context, limit *int, offset *int) ([]domain.Product, error)
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}
