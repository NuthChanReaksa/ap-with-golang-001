package main

import (
	"github.com/gin-gonic/gin"
	"homework1/internal/config"
	"homework1/internal/database"
	"homework1/internal/repository"
	"homework1/internal/routers"
	"homework1/internal/services"
	"log"
)

func main() {

	// load configuration
	config := config.LoadConfig()

	// create database connection
	db := database.InitDB(config.DatabaseDN)

	log.Println("Database connection successful", db)

	user_repository := repository.NewUserRepository(db)
	product_repository := repository.NewProductRepository(db)

	user_service := services.NewUserService(user_repository)
	product_service := services.NewProductService(product_repository)

	router := gin.Default()
	routers.SetupRouter(router, user_service, product_service)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
