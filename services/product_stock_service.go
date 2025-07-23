package services

import "github.com/kryast/Crud-7.git/repositories"

type ProductStockService interface {
}

type productStockService struct {
	repo repositories.ProductStockRepository
}

func NewProductStockService(r repositories.ProductStockRepository) ProductStockService {
	return &productStockService{r}
}
