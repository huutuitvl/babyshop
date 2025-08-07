package domain

import (
	"time"

	"gorm.io/gorm"
)

type AuditTrail struct {
	CreatedAt   time.Time
	CreatedByID *uint
	UpdatedAt   time.Time
	UpdatedByID *uint
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	DeletedByID *uint
}
