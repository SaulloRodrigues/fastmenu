package routes

import (
	"menu-delivery/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rg *gin.RouterGroup) {
	categoryGroup := rg.Group("/categories")
	{
		categoryGroup.POST("/", controllers.CreateCategory)
		categoryGroup.GET("/", controllers.GetCategories)
		categoryGroup.GET("/:id", controllers.GetCategory)
		categoryGroup.PUT("/:id", controllers.UpdateCategory)
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}
}
