package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/services"
)

type ProductStockHandler struct{ svc services.ProductStockService }

func NewProductStockHandler(s services.ProductStockService) *ProductStockHandler {
	return &ProductStockHandler{s}
}

func (h *ProductStockHandler) Create(c *gin.Context) {
	var m models.ProductStock
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, m)
}
