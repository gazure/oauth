package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gazure/oauth/token-generators"
)

const paramClientId = "client_id"
const paramClientSecret = "client_secret"
const postParamScope = "scope"
const paramGrantType = "grant_type"

const grantTypeClientCredentials = "client_credentials"

type user struct {
	name     string
	password string // TODO: not this
}

var users = map[string]user{
	"Grant": {name: "Grant", password: "987654321"},
	"Jim":   {name: "Jim", password: "password"},
}

type handler func(ctx *gin.Context)

var handlerMap = map[string]handler{
	grantTypeClientCredentials: handleClientCredentials,
}

func badRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"error": message,
	})
}

func validateClientCredentials(clientId string, clientSecret string) error {
	err := errors.New("invalid user name or password")
	for _, element := range users {
		if element.name == clientId {
			if element.password != clientSecret {
				return err
			}
			return nil
		}
	}
	return err
}

func handleClientCredentials(c *gin.Context) {
	clientId := c.PostForm(paramClientId)
	err := validateClientCredentials(clientId, c.PostForm(paramClientSecret))
	if err != nil {
		c.JSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}
	token, err := token_generators.IssueJwt(clientId, rsaCertificate)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  token,
		"refresh_token": "bar",
		"scope":         c.PostForm(postParamScope),
		"expires_in":    3600,
	})

}

func token(c *gin.Context) {
	grantType := c.PostForm(paramGrantType)
	if handlerFunction, ok := handlerMap[grantType]; ok {
		handlerFunction(c)
	} else {
		badRequest(c, "invalid grant type specified")
	}
}
