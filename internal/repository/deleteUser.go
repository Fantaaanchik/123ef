package repository

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repeatTestProject/internal/db"
	"repeatTestProject/models"
)

func (user *UserRepository) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if result := db.GetDB().Delete(&models.User{}, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user by this id does not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User is successfully deleted in table users"})
}
