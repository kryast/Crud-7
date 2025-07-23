package main

import (
	"log"

	"github.com/kryast/Crud-7.git/database"
	"github.com/kryast/Crud-7.git/handlers"
	"github.com/kryast/Crud-7.git/models"
	"github.com/kryast/Crud-7.git/repositories"
	"github.com/kryast/Crud-7.git/router"
	"github.com/kryast/Crud-7.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("DB connect fail:", err)
	}

	db.AutoMigrate(&models.Product{}, &models.Supplier{}, &models.ProductStock{})

	prodRepo := repositories.NewProductRepository(db)
	suppRepo := repositories.NewSupplierRepository(db)
	stockRepo := repositories.NewProductStockRepository(db)

	prodSvc := services.NewProductService(prodRepo)
	suppSvc := services.NewSupplierService(suppRepo)
	stockSvc := services.NewProductStockService(stockRepo)

	prodH := handlers.NewProductHandler(prodSvc)
	suppH := handlers.NewSupplierHandler(suppSvc)
	stockH := handlers.NewProductStockHandler(stockSvc)

	r := router.SetupRouter(prodH, suppH, stockH)
	log.Println("🚀 Server running at :8080")
	r.Run(":8080")
}
