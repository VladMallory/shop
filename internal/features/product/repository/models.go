package product_repository

import "authTest/internal/features/product/domain"

type ProductModel struct {
	ID          int    `db:"id"`
	Version     int    `db:"version"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
}

func productDomainFromModel(model ProductModel) domain.Product {
	return domain.Product{
		ID:          model.ID,
		Version:     model.Version,
		Title:       model.Title,
		Description: model.Description,
		Price:       model.Price,
	}
}

func productDomainFromModels(models []ProductModel) []domain.Product {
	products := make([]domain.Product, len(models))
	for i, model := range models {
		products[i] = productDomainFromModel(model)
	}

	return products
}
