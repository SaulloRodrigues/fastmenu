package controllers

import (
	"app-fastmenu-backend/models"
	"app-fastmenu-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	var products []models.Product

	result := services.DB.Find(&products)

	if result.Error != nil {
		log.Fatal("Erro ao buscar produtos: ", result.Error)
	}

	c.JSON(http.StatusAccepted, products)
}


func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler o corpo da requisição " + err.Error()})
		return
	}

	if err := services.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)

}