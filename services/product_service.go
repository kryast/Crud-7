package services

import (
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/repositories"
)

type ProductService interface {
	Create(*models.Product) error
	All() ([]models.Product, error)
	ByID(id uint) (models.Product, error)
	Update(*models.Product) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService { return &productService{r} }

func (s *productService) Create(m *models.Product) error       { return s.repo.Create(m) }
func (s *productService) All() ([]models.Product, error)       { return s.repo.FindAll() }
func (s *productService) ByID(id uint) (models.Product, error) { return s.repo.FindByID(id) }
func (s *productService) Update(m *models.Product) error       { return s.repo.Update(m) }
