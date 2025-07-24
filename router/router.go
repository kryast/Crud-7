package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-7.git/handlers"
)

func SetupRouter(ph *handlers.ProductHandler, sh *handlers.SupplierHandler, psh *handlers.ProductStockHandler) *gin.Engine {
	r := gin.Default()

	// Product routes
	r.POST("/products", ph.Create)
	r.GET("/products", ph.GetAll)
	r.GET("/products/:id", ph.GetByID)
	r.PUT("/products/:id", ph.Update)
	r.DELETE("/products/:id", ph.Delete)

	return r
}
