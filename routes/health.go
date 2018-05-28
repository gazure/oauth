package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gazure/oauth/models"
)

func health(c *gin.Context) {
	models.CreateUser("Grant", "987654321")
	models.CreateUser("Jim", "password")
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
