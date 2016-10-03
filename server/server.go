package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

func newServer() *gin.Engine {
	r := gin.New()

	r.Static("/assets", "./build")
	r.LoadHTMLFiles("./build/index.html")

	r.Use(ginrus.Ginrus(log, time.RFC3339, false))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Job Hunt",
		})
	})

	return r
}
