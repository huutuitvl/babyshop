package usecase

import (
	"babyshop/internal/domain"
	"babyshop/internal/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (u *ProductUsecase) GetAll() ([]domain.Product, error) {
	return u.repo.GetAll()
}

func (u *ProductUsecase) GetByID(id uint) (*domain.Product, error) {
	return u.repo.GetByID(id)
}

func (u *ProductUsecase) Create(p *domain.Product) error {
	return u.repo.Create(p)
}

func (u *ProductUsecase) Update(p *domain.Product) error {
	return u.repo.Update(p)
}

func (u *ProductUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
