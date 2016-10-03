package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./build")
	r.LoadHTMLFiles("./client/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Job Hunt",
		})
	})

	r.Run(":8080")
}
