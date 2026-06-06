package domain

import (
	"authTest/internal/errs"
	"fmt"
	"unicode/utf8"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Price       int
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
