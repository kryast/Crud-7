package repositories

import (
	"github.com/kryast/Crud-7.git/models"
	"gorm.io/gorm"
)

type ProductStockRepository interface {
	Create(stock *models.ProductStock) error
	FindAll() ([]models.ProductStock, error)
	FindByID(id uint) (models.ProductStock, error)
}

type productStockRepo struct{ db *gorm.DB }

func NewProductStockRepository(db *gorm.DB) ProductStockRepository { return &productStockRepo{db} }

func (r *productStockRepo) Create(stock *models.ProductStock) error { return r.db.Create(stock).Error }

func (r *productStockRepo) FindAll() ([]models.ProductStock, error) {
	var stocks []models.ProductStock
	err := r.db.Preload("Product").Preload("Supplier").Find(&stocks).Error
	return stocks, err
}

func (r *productStockRepo) FindByID(id uint) (models.ProductStock, error) {
	var stock models.ProductStock
	err := r.db.Preload("Product").Preload("Supplier").First(&stock, id).Error
	return stock, err
}
