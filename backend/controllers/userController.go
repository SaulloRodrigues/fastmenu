package controllers

import (
	"app-fastmenu-backend/models"
	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "password456"},
}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

// func GetUser(c *gin.Context) {
// 	user := c.Param(("id"))
// }
