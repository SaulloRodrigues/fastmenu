package main

import (
	"app-fastmenu-backend/models"
	"app-fastmenu-backend/routes"
	"app-fastmenu-backend/config"
	"os"
)

func main() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{})
	port := os.Getenv("APP_PORT")
	r := routes.SetupRouter()
	r.Run(":" + port)
}
