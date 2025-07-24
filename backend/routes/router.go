package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	
	UserRoutes(api)
	ProductRoutes(api)

	return r
}