package services

import (
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/repositories"
)

type SupplierService interface {
	Create(*models.Supplier) error
}

type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(r repositories.SupplierRepository) SupplierService {
	return &supplierService{r}
}

func (s *supplierService) Create(m *models.Supplier) error { return s.repo.Create(m) }
