package domain

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	Password     string
	Role         string     // "admin", "staff"
	Status       int        // 1: new, 2: active, 3: blocked
	VerifiedAt   *time.Time `gorm:"type:datetime"` // yyyy-mm-dd H:i:s
	VerifiedByID *uint      // admin ID
	
	AuditTrail
}
