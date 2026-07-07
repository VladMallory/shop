package product_repository

import (
	"authTest/internal/errs"
	"authTest/internal/features/product/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *ProductRepository) GetProduct(ctx context.Context, id int) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, price
	FROM products
	WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)
	var productModel ProductModel

	err := row.Scan(
		&productModel.ID,
		&productModel.Version,
		&productModel.Title,
		&productModel.Description,
		&productModel.Price,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, fmt.Errorf("product not found: %w", errs.ErrNotFound)
		}

		return domain.Product{}, fmt.Errorf("scan error: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
