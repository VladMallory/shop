package product_service

import (
	"authTest/internal/features/product/domain"
	"context"
	"fmt"
)

func (s *ProductService) PatchProduct(ctx context.Context, id int, patch domain.ProductPatch) (domain.Product, error) {
	product, err := s.productRepository.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	if err := product.ApplyPatch(patch); err != nil {
		return domain.Product{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	productFromRepo, err := s.productRepository.PatchProduct(ctx, id, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to patch product: %w", err)
	}

	return productFromRepo, nil
}
