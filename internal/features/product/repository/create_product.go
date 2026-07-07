package product_repository

import (
	"authTest/internal/features/product/domain"
	"context"
	"fmt"
)

func (r *ProductRepository) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO products (title, description, price)
	VALUES($1, $2, $3)
	RETURNING id, version, title, description, price;
	`

	var productModel ProductModel
	row := r.pool.QueryRow(ctx, query, product.Title, product.Description, product.Price)

	err := row.Scan(
		&productModel.ID,
		&productModel.Version,
		&productModel.Title,
		&productModel.Description,
		&productModel.Price,
	)

	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to create product: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
