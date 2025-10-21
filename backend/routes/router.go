package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.HEAD("/health", func(c *gin.Context) {
		c.Status(200)
	})

	api := r.Group("/api")

	UserRoutes(api)
	ProductRoutes(api)
	CategoryRoutes(api)

	return r
}
