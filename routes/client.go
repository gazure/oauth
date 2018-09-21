package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gazure/oauth/models"
	"net/http"
	"github.com/satori/go.uuid"
)

const paramClientName = "client_name"
const paramClientDescription = "description"

func registerNewClient(c *gin.Context) {
	clientName := c.PostForm(paramClientName)
	clientDescription := c.PostForm(paramClientDescription)
	client := models.CreateClient(clientName, clientDescription)
	c.JSON(http.StatusOK, client.ToDTO())
}


func listClients(c *gin.Context) {
	clients := models.GetAllClients()
	dtos := make([]gin.H, len(clients))
	for i, client := range clients {
		dtos[i] = client.ToDTO()
	}
	response := gin.H{
		"clients": dtos,
	}
	c.JSON(http.StatusOK, response)
}

func generateNewSecret(c *gin.Context) {
	id, err := uuid.FromString(c.Query("client_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad UUID",
		})
		return
	}

	client := models.GenerateNewSecret(id)
	response := gin.H{
		"Client": client.ToDTO(),
	}
	c.JSON(http.StatusOK, response)
}
