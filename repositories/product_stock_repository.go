package repositories

import (
	"github.com/kryast/Crud-7.git/models"
	"gorm.io/gorm"
)

type ProductStockRepository interface {
	Create(stock *models.ProductStock) error
}

type productStockRepo struct{ db *gorm.DB }

func NewProductStockRepository(db *gorm.DB) ProductStockRepository { return &productStockRepo{db} }

func (r *productStockRepo) Create(stock *models.ProductStock) error { return r.db.Create(stock).Error }
