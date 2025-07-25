package handlers

import (
	"net/http"
	"strconv"

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

func (h *ProductStockHandler) GetAll(c *gin.Context) {
	ms, _ := h.svc.All()
	c.JSON(http.StatusOK, ms)
}

func (h *ProductStockHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m, err := h.svc.ByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, m)
}

func (h *ProductStockHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m models.ProductStock
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m.ID = uint(id)
	if err := h.svc.Update(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, m)
}
