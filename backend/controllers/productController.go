package controllers

import (
	"app-fastmenu-backend/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := []models.Product{
		{ID: 1, Name: "Pão Francês", Price: 1.5, Description: "Pão francês fresquinho"},
		{ID: 2, Name: "Bolo de Milho", Price: 5.0, Description: "Bolo de milho cremoso"},
	}
	c.JSON(200, products)
}