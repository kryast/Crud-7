package repositories

import (
	"github.com/kryast/Crud-7.git/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (models.Product, error)
}

type productRepo struct{ db *gorm.DB }

func NewProductRepository(db *gorm.DB) ProductRepository { return &productRepo{db} }

func (r *productRepo) Create(product *models.Product) error { return r.db.Create(product).Error }

func (r *productRepo) FindAll() ([]models.Product, error) {
	var products []models.Product
	return products, r.db.Find(&products).Error
}

func (r *productRepo) FindByID(id uint) (models.Product, error) {
	var product models.Product
	return product, r.db.First(&product, id).Error
}
