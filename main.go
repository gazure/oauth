package main

import (
	"github.com/gazure/oauth/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gazure/oauth/models"
)

func main() {
	models.Init()
	models.Migrate()
	routes.LoadCertificate()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	routes.ConfigureRoutes(r)
	r.Run()
}
