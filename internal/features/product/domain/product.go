package domain

import (
	"authTest/internal/errs"
	"fmt"
	"unicode/utf8"
)

const (
	productIDUninitialized      = -1
	productVersionUninitialized = -1
)

type Product struct {
	ID          int
	Version     int
	Title       string
	Description string
	Price       int
}

type ProductPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Price       Nullable[int]
}

func NewProductPatch(title Nullable[string], description Nullable[string], price Nullable[int]) ProductPatch {
	return ProductPatch{
		Title:       title,
		Description: description,
		Price:       price,
	}
}

func (p *ProductPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("title cannot be set to NULL: %w", errs.ErrInvalidArgument)
	}

	if p.Description.Set && p.Description.Value == nil {
		return fmt.Errorf("description cannot be set to NULL: %w", errs.ErrInvalidArgument)
	}

	if p.Price.Set && p.Price.Value == nil {
		return fmt.Errorf("price cannot be set to NULL: %w", errs.ErrInvalidArgument)
	}

	return nil
}

func (p *Product) ApplyPatch(patch ProductPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("product patch is invalid: %w", err)
	}

	tmp := *p

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = *patch.Description.Value
	}

	if patch.Price.Set {
		tmp.Price = *patch.Price.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("patched product is invalid: %w", err)
	}

	*p = tmp

	return nil
}

func NewUninitializedProduct(title string, description string, price int) Product {
	return Product{
		ID:          productIDUninitialized,
		Version:     productVersionUninitialized,
		Title:       title,
		Description: description,
		Price:       price,
	}
}

func (p *Product) Validate() error {
	titleLength := utf8.RuneCountInString(p.Title)
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf("title length has to be between 1 and 100: %w", errs.ErrInvalidArgument)
	}

	descriptionLength := utf8.RuneCountInString(p.Description)
	if descriptionLength < 1 || descriptionLength > 1000 {
		return fmt.Errorf(
			"description length has to be between 1 and 1000: %w",
			errs.ErrInvalidArgument,
		)
	}

	if p.Price <= 0 {
		return fmt.Errorf("price has to be positive: %w", errs.ErrInvalidArgument)
	}

	return nil
}
