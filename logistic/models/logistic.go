package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base struct
type Base struct {
	ID        int64      `gorm:"autoIncrement"`
	UUID      string     `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate creates a UUID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New().String()
	return
}

// Logistic struct
type Logistic struct {
	Base
	LogisticName    string `gorm:"not null"`
	Amount          int64  `gorm:"not null"`
	DestinationName string `gorm:"not null"`
	OriginName      string `gorm:"not null"`
	Duration        string `gorm:"not null"`
}
