package main

import (
	"app-fastmenu-backend/routes"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	r := routes.SetupRouter()
	r.Run(":" + port)
}
