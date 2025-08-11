package controllers

import (
	"app-fastmenu-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		log.Println("Erro ao buscar produtos: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	product, err := services.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
    var input services.ProductInputDetails

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler corpo da requisição: " + err.Error()})
        return
    }

    product, err := services.CreateProduct(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto: " + err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var input services.ProductInputDetails

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler o corpo da requisição: " + err.Error()})
		return
	}

	if err := services.UpdateProductById(id, input); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado ou não foi possível atualizar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso!"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteProductById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado ou não foi possível deletar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso!"})
}
