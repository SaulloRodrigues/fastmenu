package main

import (
	"menu-delivery/config"
	"menu-delivery/models"
	"menu-delivery/routes"
	"os"
)

func main() {
	config.ConnectToDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{})
	port := os.Getenv("APP_PORT")
	r := routes.SetupRouter()
	r.Run(":" + port)
}