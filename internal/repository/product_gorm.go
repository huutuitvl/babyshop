package repository

import (
	"babyshop/internal/domain"

	"gorm.io/gorm"
)

type productGormRepo struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of productGormRepo with the given gorm.DB.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productGormRepo{db}
}

// GetAll retrieves all products from the database.
// Returns a slice of domain.Product and an error if the operation fails.
func (r *productGormRepo) GetAll() ([]domain.Product, error) {
	var products []domain.Product

	// Select only the necessary fields to optimize performance
	// This avoids loading unnecessary data from the database
	err := r.db.Select("id, name, description, price, size, stock, image_url").Find(&products).Error

	return products, err
}

// GetByID retrieves a product from the database by its unique ID.
// Returns a pointer to domain.Product and an error if the operation fails.
func (r *productGormRepo) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Select("id, name, description, price, size, stock, image_url").First(&product, id).Error

	return &product, err
}

// Create inserts a new product into the database.
// Returns an error if the operation fails.
func (r *productGormRepo) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

// Update modifies an existing product in the database.
// Returns an error if the operation fails.
func (r *productGormRepo) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

// Delete removes a product from the database by its unique ID.
// Returns an error if the operation fails.
func (r *productGormRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
