package product_transport

import "authTest/internal/domain"

type ProductDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func productDTOFromDomain(product domain.Product) ProductDTO {
	return ProductDTO{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
	}
}

func productsDTOFromDomains(products []domain.Product) []ProductDTO {
	productDTOs := make([]ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = productDTOFromDomain(product)
	}

	return productDTOs
}
