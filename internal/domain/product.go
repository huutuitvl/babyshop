package domain

type Product struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Slug     string `gorm:"type:varchar(191);uniqueIndex"` // SEO URL
	Price    float64
	Size     string
	Stock    int
	ImageURL string // URL to product image

	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	OGImage         string // URL to product image when shared
	CategoryID      uint
	Category        Category `gorm:"foreignKey:CategoryID"`

	AuditTrail
}
