package domain

type Category struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Slug        string `gorm:"uniqueIndex"`
	Description string
	Products    []Product

	AuditTrail
}
