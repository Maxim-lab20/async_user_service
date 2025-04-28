package controller

import (
	"async_user_service/app/dto"
	"async_user_service/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type UserController struct {
	userService *service.UserService
}

var (
	userControllerInstance *UserController
	once                   sync.Once
)

func GetUserController() *UserController {
	once.Do(func() {
		userControllerInstance = &UserController{
			userService: service.NewUserService(),
		}
	})
	return userControllerInstance
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userService.CreateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
