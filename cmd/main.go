package main

import (
	"babyshop/internal/config"
	"babyshop/internal/delivery/http"
	"babyshop/internal/domain"
	"babyshop/internal/repository"
	"babyshop/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db := config.ConnectDB()

	// Create a new Gin router
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Automatically migrate the Product model (creates table if not exists)
	db.AutoMigrate(&domain.User{}, &domain.Category{}, &domain.Product{})
	db.AutoMigrate(&domain.Order{}, &domain.OrderItem{})
	db.AutoMigrate(&domain.ProductImage{})

	config.SeedProducts(db)
	config.SeedAdminUser(db)
	config.SeedDummyOrders(db)

	// Register authentication HTTP handlers/routes
	http.NewAuthHandler(r, db)
	http.NewAdminHandler(r, db)

	// Set up repository and usecase for Product
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)
	// Register product HTTP handlers/routes
	http.NewProductHandler(r, productUC)

	userUC := usecase.NewUserUsecase(db)
	http.NewUserHandler(r, userUC)

	// Start the HTTP server on port 8080
	r.Run(":8080")
}
