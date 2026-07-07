package product_repository

import (
	"authTest/internal/features/product/domain"
	"context"
	"fmt"
)

func (r *ProductRepository) GetProducts(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, price
	FROM products
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2
    `

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to select products: %w", err)
	}

	defer rows.Close()
	var productModels []ProductModel
	for rows.Next() {
		var productModel ProductModel
		if err := rows.Scan(
			&productModel.ID,
			&productModel.Version,
			&productModel.Title,
			&productModel.Description,
			&productModel.Price,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		productModels = append(productModels, productModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows next error: %w", err)
	}

	return productDomainFromModels(productModels), nil
}
