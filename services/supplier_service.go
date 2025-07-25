package services

import (
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/repositories"
)

type SupplierService interface {
	Create(*models.Supplier) error
	All() ([]models.Supplier, error)
	ByID(id uint) (models.Supplier, error)
	Update(*models.Supplier) error
}

type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(r repositories.SupplierRepository) SupplierService {
	return &supplierService{r}
}

func (s *supplierService) Create(m *models.Supplier) error       { return s.repo.Create(m) }
func (s *supplierService) All() ([]models.Supplier, error)       { return s.repo.FindAll() }
func (s *supplierService) ByID(id uint) (models.Supplier, error) { return s.repo.FindByID(id) }
func (s *supplierService) Update(m *models.Supplier) error       { return s.repo.Update(m) }
