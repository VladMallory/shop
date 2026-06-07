package product_service

import (
	"authTest/internal/features/product/domain"
	"context"
)

func (s *ProductService) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := product.Validate(); err != nil {
		return domain.Product{}, err
	}

	product, err := s.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
