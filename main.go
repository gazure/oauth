package main

import "github.com/gin-gonic/gin"

const POST_PARAM_CLIENT_ID = "client_id"
const POST_PARAM_CLIENT_SECRET = "client_secret"
const POST_PARAM_SCOPE = "scope"
const POST_PARAM_GRANT_TYPE = "grant_type"


type user struct {
	name string
	password string // TODO: not this
}


var USERS = []user{
	{name: "Grant", password: "987654321"},
	{name: "Jim", password: "password"},
}


func bad_request(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "no grant_type specified",
	})
}

func token(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	if grantType == "" {
		bad_request(c)
		return
	}

}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func main() {
	r := gin.Default()
	r.POST("/token", token)
	r.GET("/health", health)
	r.Run()
}
