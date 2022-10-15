package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"localhost.com/game_application"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/", game_application.Game)
	router.Run("0.0.0.0:8081")
}
