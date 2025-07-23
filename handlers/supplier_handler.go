package handlers

import "github.com/kryast/Crud-7.git/services"

type SupplierHandler struct{ svc services.SupplierService }

func NewSupplierHandler(s services.SupplierService) *SupplierHandler { return &SupplierHandler{s} }
