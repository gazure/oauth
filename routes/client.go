package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gazure/oauth/models"
	"net/http"
)

const paramClientName = "client_name"
const paramClientDescription = "description"

func registerNewClient(c *gin.Context) {
	clientName := c.PostForm(paramClientName)
	clientDescription := c.PostForm(paramClientDescription)
	client := models.CreateClient(clientName, clientDescription)
	c.JSON(http.StatusOK, client.ToDTO())
}
