package repositories

import (
	"gorm.io/gorm"
)

type SupplierRepository interface {
}

type supplierRepo struct{ db *gorm.DB }

func NewSupplierRepository(db *gorm.DB) SupplierRepository { return &supplierRepo{db} }
