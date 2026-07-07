package product_repository

import (
	"authTest/internal/errs"
	"authTest/internal/features/product/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *ProductRepository) PatchProduct(ctx context.Context, id int, patched_product domain.Product) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE product
	SET 
	    title=$1
		description=$2
		price=$3
		version = version+1
	WHERE id=$4 AND version=$5
	RETURNING 
		id, version, title, description, price;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		patched_product.Title,
		patched_product.Description,
		patched_product.Price,
		patched_product.Version,
		patched_product.Version,
	)

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
			return domain.Product{}, fmt.Errorf("product concurrently accessed: %v: %w", err, errs.ErrConflict)
		}

		return domain.Product{}, fmt.Errorf("scan error: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
