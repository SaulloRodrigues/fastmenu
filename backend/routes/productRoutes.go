package routes 

import (
	"app-fastmenu-backend/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	productGroup := rg.Group("/products")
	{
		productGroup.GET("/", controllers.GetProducts)
	}
}