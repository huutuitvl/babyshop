package config

import (
	"babyshop/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	products := []domain.Product{
		{Name: "Áo thun bé trai", Price: 120000, Size: "M", Stock: 10, ImageURL: "https://cdn.example.com/aothun.jpg"},
		{Name: "Váy công chúa", Price: 180000, Size: "S", Stock: 5, ImageURL: "https://cdn.example.com/vaycongchua.jpg"},
		{Name: "Áo khoác lông", Price: 250000, Size: "L", Stock: 7, ImageURL: "https://cdn.example.com/aokhoac.jpg"},
		{Name: "Quần jean bé gái", Price: 160000, Size: "M", Stock: 12, ImageURL: "https://cdn.example.com/quanjean.jpg"},
		{Name: "Bộ đồ ngủ trẻ em", Price: 90000, Size: "S", Stock: 20, ImageURL: "https://cdn.example.com/dongu.jpg"},
	}

	for _, p := range products {
		db.FirstOrCreate(&p, domain.Product{Name: p.Name}) // tránh trùng
	}
}

func SeedAdminUser(db *gorm.DB) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("secret123"), bcrypt.DefaultCost)
	admin := domain.User{
		Email:    "admin@example.com",
		Password: string(hashed),
		Role:     "admin",
	}
	db.FirstOrCreate(&admin, domain.User{Email: admin.Email})
}

func SeedDummyOrders(db *gorm.DB) {
	var user domain.User
	db.First(&user)

	var p1, p2 domain.Product
	db.First(&p1)
	db.Offset(1).First(&p2)

	order := domain.Order{
		UserID:     user.ID,
		TotalPrice: p1.Price + p2.Price,
		Status:     "pending",
		Items: []domain.OrderItem{
			{ProductID: p1.ID, Qty: 1, UnitPrice: p1.Price},
			{ProductID: p2.ID, Qty: 1, UnitPrice: p2.Price},
		},
	}

	db.Create(&order)
}
