package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductStock struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ProductID   uint           `json:"product_id" binding:"required"`
	Product     Product        `json:"product" gorm:"foreignKey:ProductID"`
	SupplierID  uint           `json:"supplier_id" binding:"required"`
	Supplier    Supplier       `json:"supplier" gorm:"foreignKey:SupplierID"`
	Quantity    int            `json:"quantity" binding:"required,min=0"`
	LastRestock *time.Time     `json:"last_restock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
