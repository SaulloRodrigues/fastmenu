package controllers

import (
	"log"
	"menu-delivery/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		log.Println("Erro ao buscar usuarios: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Erro ao Buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var input services.UserInputDetails

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler corpo da requisição: " + err.Error()})
		return
	}

	user, err := services.CreateUser(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
