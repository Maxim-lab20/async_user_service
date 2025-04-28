package main

import (
	"async_user_service/app/config"
	"async_user_service/app/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	context := gin.Default()
	config.GetDBConnection()
	config.GetRedisClient()
	router.Init(context)

	port := os.Getenv("PORT")
	context.Run(":" + port)
}
