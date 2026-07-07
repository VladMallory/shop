package product_service

import (
	"authTest/internal/errs"
	"authTest/internal/features/product/domain"
	"context"
	"fmt"
)

func (s *ProductService) GetProduct(ctx context.Context, id int) (domain.Product, error) {
	if id < 1 {
		return domain.Product{}, fmt.Errorf("product id cannot be less than 1: %w", errs.ErrInvalidArgument)
	}

	product, err := s.productRepository.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to get product from repo: %w", err)
	}

	return product, nil
}
