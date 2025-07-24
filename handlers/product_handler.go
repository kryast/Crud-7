package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/services"
)

type ProductHandler struct{ svc services.ProductService }

func NewProductHandler(s services.ProductService) *ProductHandler { return &ProductHandler{s} }

func (h *ProductHandler) Create(c *gin.Context) {
	var m models.Product
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
