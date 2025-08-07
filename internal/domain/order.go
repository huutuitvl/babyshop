package domain

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint
	User       User        `gorm:"foreignKey:UserID"` // 👈 THÊM dòng này để dùng order.User.Name
	TotalPrice float64
	Status     string
	Items      []OrderItem `gorm:"foreignKey:OrderID"`

	AuditTrail
}
