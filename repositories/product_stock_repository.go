package repositories

import (
	"gorm.io/gorm"
)

type ProductStockRepository interface {
}

type productStockRepo struct{ db *gorm.DB }

func NewProductStockRepository(db *gorm.DB) ProductStockRepository { return &productStockRepo{db} }
