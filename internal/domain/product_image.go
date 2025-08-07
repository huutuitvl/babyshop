package domain

type ProductImage struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint
	URL       string
	IsDefault bool // ảnh chính?

	Product Product

	AuditTrail
}
