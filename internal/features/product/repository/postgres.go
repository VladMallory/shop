package product_repository

import database_postgres "authTest/internal/platform/db"

type ProductRepository struct {
	pool *database_postgres.Pool
}

func NewProductRepository(pool *database_postgres.Pool) *ProductRepository {
	return &ProductRepository{
		pool: pool,
	}
}
