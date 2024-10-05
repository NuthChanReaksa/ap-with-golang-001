
package routers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"homework1/internal/services"	
	models "homework1/internal/models"
)


// Define route handlers
func GetAllUsers(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := (userService).GetAllUsers();
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := uuid.Parse(c.Param("id"))
		user, err := (userService).GetUserById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdUser, err := (userService).CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, createdUser)
	}
}

func UpdateUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := uuid.Parse(c.Param("id"))
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedUser, err := (userService).UpdateUser(id, user)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := uuid.Parse(c.Param("id"))
		err := (userService).DeleteUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
