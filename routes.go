package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	// health-check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// API v1
	{
		v1 := router.Group("/api/v1")

		v1.GET("/", func(c *gin.Context) {})
	}
}
