package repository

import (
	"babyshop/internal/domain"
)

// ProductRepository defines the interface for product data operations.
// It provides methods to retrieve, create, update, and delete products.
//
// GetAll retrieves all products from the repository.
// GetByID retrieves a product by its unique identifier.
// Create adds a new product to the repository.
// Update modifies an existing product in the repository.
// Delete removes a product from the repository by its unique identifier.
type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id uint) error
}
