package services

import "github.com/kryast/Crud-7.git/repositories"

type SupplierService interface {
}

type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(r repositories.SupplierRepository) SupplierService {
	return &supplierService{r}
}
