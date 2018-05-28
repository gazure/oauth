package main

import (
	"github.com/gazure/oauth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	routes.LoadCertificate()
	r := gin.Default()
	routes.ConfigureRoutes(r)
	r.Run()
}
