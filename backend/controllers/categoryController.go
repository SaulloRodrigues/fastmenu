package controllers

import (
	"app-fastmenu-backend/services"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		log.Println("Erro ao buscar categorias: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	id := c.Params.ByName("id")

	category, err := services.GetCategoryById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
		return
	}
	
	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var input services.CategoryInputDetails

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler corpo da requisição: " + err.Error()})
		return
	}

	category, err := services.CreateCategory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar categoria: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var input services.CategoryInputDetails

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler o corpo da requisição: " + err.Error()})
		return
	}

	if err := services.UpdateCategoryById(id, input); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada ou não foi possível atualizar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria atualizada com sucesso!"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteCategoryById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada ou não foi possível deletar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria deletada com sucesso!"})
}
