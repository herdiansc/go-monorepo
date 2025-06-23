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

// Auth struct
type Auth struct {
	Base
	MSISDN   string `gorm:"size:16;not null;unique"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Name     string
}
