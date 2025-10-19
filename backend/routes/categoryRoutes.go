package routes

import (
	"app-fastmenu-backend/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/api/categories")
	{
		categoryGroup.POST("/", controllers.CreateCategory)
		categoryGroup.GET("/", controllers.GetCategories)
		categoryGroup.GET("/:id", controllers.GetCategory)
		categoryGroup.PUT("/:id", controllers.UpdateCategory)
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}
}