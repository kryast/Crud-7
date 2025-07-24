package repositories

import (
	"github.com/kryast/Crud-7.git/models"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *models.Supplier) error
}

type supplierRepo struct{ db *gorm.DB }

func NewSupplierRepository(db *gorm.DB) SupplierRepository { return &supplierRepo{db} }

func (r *supplierRepo) Create(supplier *models.Supplier) error { return r.db.Create(supplier).Error }
