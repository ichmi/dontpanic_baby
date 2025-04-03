package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ichmi/dpbaby/srcs/game"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/styles.css", "./assets/styles.css")
	router.StaticFile("/star.png", "./assets/star.png")
	router.StaticFile("/main.js", "./assets/main.js")
	router.LoadHTMLGlob("assets/*")

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		apiURL := os.Getenv("API_URL")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"API_URL": apiURL,
		})
	})
	router.POST("/", game.ResponseGame)

	router.Run(":8080")
}
