package domain

type OrderItem struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"` // 👈 để gọi item.Product.Name
	Price     float64
	Quantity  int
	Subtotal  float64
}

