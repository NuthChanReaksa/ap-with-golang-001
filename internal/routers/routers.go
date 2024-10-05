package routers


import (
	"github.com/gin-gonic/gin"
	"homework1/internal/services"
)

func SetupRouter(router *gin.Engine, userService services.UserService, productService services.ProductService) {
	// User routes
	userGroup := router.Group("/users")
	{
		userGroup.GET("", GetAllUsers(userService))
		userGroup.GET("/:id", GetUser(userService))
		userGroup.POST("", CreateUser(userService))
		userGroup.PUT("/:id", UpdateUser(userService))
		userGroup.DELETE("/:id", DeleteUser(userService))
	}

	 // Product routes
	productGroup := router.Group("/products")
	{
		productGroup.GET("", GetAllProducts(productService))
		productGroup.GET("/:id", GetProduct(productService))
		productGroup.POST("", CreateProduct(productService))
		productGroup.PUT("/:id", UpdateProduct(productService))
		productGroup.DELETE("/:id", DeleteProduct(productService))
	}
}
