package routes

import (
	"app-fastmenu-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
	}
}