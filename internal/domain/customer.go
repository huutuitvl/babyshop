package domain

type Customer struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Email   string
	Phone   string
	Address string
	Status  string // pending, active, blocked

	AuditTrail
}
