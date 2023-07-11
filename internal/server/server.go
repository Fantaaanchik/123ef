package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repeatTestProject/models"
)

type Handler struct {
	Engine   *gin.Engine
	services ServiceInterface
}

type ServiceInterface interface {
	GetUserFromDB() ([]models.User, error)
	AddNewUserToDB(user models.User) error
	UpdateUserDataFromDB(userID string, user models.User) error
	DeleteUserDataFromDB(userID string) error
}

func NewHandler(services ServiceInterface, engine *gin.Engine) *Handler {
	return &Handler{services: services, Engine: engine}
}

func (h Handler) AllRoutes() {
	h.Engine.GET("/get_users", h.GetUsers)
	h.Engine.POST("/add_users", h.AddUsers)
	h.Engine.PUT("/update_user_data/:id", h.UpdateUserData)
	h.Engine.DELETE("/delete_user_data/:id", h.DeleteUserData)
}

func (h Handler) GetUsers(c *gin.Context) {
	user, err := h.services.GetUserFromDB()
	if err != nil {
		log.Println("err: ", err)
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h Handler) AddUsers(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := h.services.AddNewUserToDB(user)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Новый пользователь добавлен"})
}

func (h Handler) UpdateUserData(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data type"})
	}
	if err := h.services.UpdateUserDataFromDB(userID, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid in updating user data"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User data successfully updated"})
}

func (h Handler) DeleteUserData(c *gin.Context) {
	userID := c.Param("id")
	if err := h.services.DeleteUserDataFromDB(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
