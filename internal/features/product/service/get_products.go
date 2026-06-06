package product_service

import (
	"authTest/internal/domain"
	"authTest/internal/errs"
	"context"
	"fmt"
)

func (s *ProductService) GetProducts(ctx context.Context, limit *int, offset *int) ([]domain.Product, error) {
	if limit != nil && *limit < 1 {
		return nil, fmt.Errorf("limit cannot be below 1: %w", errs.ErrInvalidArgument)
	}

	if offset != nil && *offset < 1 {
		return nil, fmt.Errorf("offset cannot be below 1: %w", errs.ErrInvalidArgument)
	}

	products, err := s.productRepository.GetProducts(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}
