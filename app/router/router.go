package router

import (
	"async_user_service/app/controller"
	"github.com/gin-gonic/gin"
)

func Init(context *gin.Engine) {

	usersGroupV1 := context.Group("/v1/users")
	{
		usersGroupV1.GET("", func(c *gin.Context) {
			controller.GetUserController().GetUsers(c)
		})
		usersGroupV1.POST("", func(c *gin.Context) {
			controller.GetUserController().CreateUser(c)

		})
	}
}
