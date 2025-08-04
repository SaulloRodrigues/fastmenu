package controllers

import (
	"app-fastmenu-backend/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := []models.Product{
		{ID: 1, Name: "Pão Francês", Price: 0.80, Description: "Pão francês fresquinho", ImageURL: "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcSjcdwMgD4mexlwQF2coR_omGVE3OPBUEGc85QKlipTkYgS5utAd1_QK9sVyxyK"},
		{ID: 2, Name: "Bolo de Milho", Price: 7.00, Description: "Bolo de milho cremoso", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRwEL8k-SsXOkSlZvQeThQnDWoTXau8UxPdxJ6gpcKF2hjM5hojWz4RqlNJlkPF"},
		{ID: 3, Name: "Coxinha", Price: 5.00, Description: "Coxinha de frango com catupiry", ImageURL: "https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcRAgUo8e0gg9jN5oR8yIa_UAcVXFim_pja1e4IjXr7X1q7JKqnhdr-BIODo5xT9"},
		{ID: 4, Name: "Pastel de Queijo", Price: 6.00, Description: "Pastel de queijo crocante", ImageURL: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcTih2i_o4wm3D9D3yWePBJlxtqNiFGKM3q3uNV8AMXvs0qQuxmHT8jRyWZfFdXM"},
		{ID: 5, Name: "Suco Natural", Price: 4.50, Description: "Suco natural de laranja", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTQ8fxTh07DBlmepYI1CCxEVsfRpEDffugs5dimBetmhzqu0AKLM9wo8pIO2gcp"},
		{ID: 6, Name: "Sanduíche Natural", Price: 8.00, Description: "Sanduíche natural com peito de peru", ImageURL: "https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcQRiORt8JlJQEWKt3rmbcblVmC78trVKWsf2cF57bqe3kQVfAOlUMOrA2jXTnox"},
	}
	c.JSON(200, products)
}