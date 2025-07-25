package services

import (
	"time"

	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/repositories"
)

type ProductStockService interface {
	Create(*models.ProductStock) error
	All() ([]models.ProductStock, error)
	ByID(id uint) (models.ProductStock, error)
}

type productStockService struct {
	repo repositories.ProductStockRepository
}

func NewProductStockService(r repositories.ProductStockRepository) ProductStockService {
	return &productStockService{r}
}

func (s *productStockService) Create(m *models.ProductStock) error {
	if m.LastRestock == nil {
		now := time.Now()
		m.LastRestock = &now
	}
	return s.repo.Create(m)
}

func (s *productStockService) All() ([]models.ProductStock, error) { return s.repo.FindAll() }

func (s *productStockService) ByID(id uint) (models.ProductStock, error) { return s.repo.FindByID(id) }
