package main

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const paramClientId = "client_id"
const paramClientSecret = "client_secret"
const postParamScope = "scope"
const paramGrantType = "grant_type"

const GrantTypeClientCredentials = "client_credentials"

type user struct {
	name     string
	password string // TODO: not this
}

var USERS = []user{
	{name: "Grant", password: "987654321"},
	{name: "Jim", password: "password"},
}

func validateClientCredentials(clientId string, clientSecret string) error {
	err := errors.New("invalid user name or password")
	for _, element := range USERS {
		if element.name == clientId {
			if element.password != clientSecret {
				return err
			}
			return nil
		}
	}
	return err
}

func badRequest(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "no grant_type specified",
	})
}

func handleClientCredentials(c *gin.Context) {
	err := validateClientCredentials(c.PostForm(paramClientId), c.PostForm(paramClientSecret))
	if err != nil {
		c.JSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  "foo",
		"refresh_token": "bar",
		"scope":         c.PostForm(postParamScope),
		"expires_in":    "never",
	})

}

func token(c *gin.Context) {
	grantType := c.PostForm(paramGrantType)
	if grantType == "" {
		badRequest(c)
		return
	}
	if grantType == GrantTypeClientCredentials {
		handleClientCredentials(c)
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
