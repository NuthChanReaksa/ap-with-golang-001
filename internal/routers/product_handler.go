package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"homework1/internal/services"
	models "homework1/internal/models" // Import the package that defines the Product type
)

// Define product handlers with improved error handling

func GetAllProducts(productService services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := productService.GetAllProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}


func GetProduct(productService services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		product, err := productService.GetProductById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(productService services.ProductService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var product models.Product
        // Bind the JSON input to the model.Product struct
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        createdProduct, err := productService.CreateProduct(product)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, createdProduct)
    }
}


func UpdateProduct(productService services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		var product models.Product
		
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedProduct, err := productService.UpdateProduct(id, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedProduct)
	}
}

func DeleteProduct(productService services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		err = productService.DeleteProduct(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
