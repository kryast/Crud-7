package repositories

import (
	"github.com/kryast/Crud-7.git/models"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *models.Supplier) error
	FindAll() ([]models.Supplier, error)
	FindByID(id uint) (models.Supplier, error)
}

type supplierRepo struct{ db *gorm.DB }

func NewSupplierRepository(db *gorm.DB) SupplierRepository { return &supplierRepo{db} }

func (r *supplierRepo) Create(supplier *models.Supplier) error { return r.db.Create(supplier).Error }

func (r *supplierRepo) FindAll() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	return suppliers, r.db.Find(&suppliers).Error
}

func (r *supplierRepo) FindByID(id uint) (models.Supplier, error) {
	var supplier models.Supplier
	return supplier, r.db.First(&supplier, id).Error
}
