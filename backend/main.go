package main

import (
	"app-fastmenu-backend/models"
	"app-fastmenu-backend/routes"
	"app-fastmenu-backend/services"
	"os"
)

func main() {
	services.ConnectToDB()
	services.DB.AutoMigrate(&models.User{}, &models.Product{})
	port := os.Getenv("APP_PORT")
	r := routes.SetupRouter()
	r.Run(":" + port)
}
