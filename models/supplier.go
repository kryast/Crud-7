package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" binding:"required,min=3"`
	Email     string         `json:"email" binding:"required,email"`
	Phone     string         `json:"phone" binding:"required"`
	Address   string         `json:"address" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
