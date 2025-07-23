package handlers

import "github.com/kryast/Crud-7.git/services"

type ProductStockHandler struct{ svc services.ProductStockService }

func NewProductStockHandler(s services.ProductStockService) *ProductStockHandler {
	return &ProductStockHandler{s}
}
