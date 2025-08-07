package usecase

import (
	"babyshop/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db *gorm.DB
}

func NewUserUsecase(db *gorm.DB) *UserUsecase {
	return &UserUsecase{db}
}

func (u *UserUsecase) Register(email, password string) error {
	var count int64
	u.db.Model(&domain.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return gorm.ErrDuplicatedKey
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := domain.User{
		Email:    email,
		Password: string(hashed),
		Role:     "staff",
		Status:   1, // New user status
	}

	return u.db.Create(&user).Error
}
