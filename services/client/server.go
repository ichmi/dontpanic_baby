package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFile("/styles.css", "./assets/styles.css")
	r.StaticFile("/star.png", "./assets/star.png")
	r.StaticFile("/main.js", "./assets/main.js")
	r.LoadHTMLGlob("assets/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":8080")
}
