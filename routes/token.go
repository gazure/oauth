package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gazure/oauth/token-generators"
	"github.com/gazure/oauth/models"
)

const paramClientId = "client_id"
const paramClientSecret = "client_secret"
const postParamScope = "scope"
const paramGrantType = "grant_type"

const grantTypeClientCredentials = "client_credentials"

var handlerMap = map[string]gin.HandlerFunc{
	grantTypeClientCredentials: handleClientCredentials,
}

func badRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"error": message,
	})
}

func validateClientCredentials(clientId string, clientSecret string) (*models.User, error) {
	err := errors.New("invalid user name or password")
	user := models.GetUser(clientId)
	if (&user).PasswordMatch(clientSecret) {
		return &user, nil
	}
	return nil, err
}

func handleClientCredentials(c *gin.Context) {
	clientId := c.PostForm(paramClientId)
	clientSecret := c.PostForm(paramClientSecret)
	userPtr, err := validateClientCredentials(clientId, clientSecret)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}
	token, err := token_generators.IssueJwt(userPtr.GetId(), rsaCertificate)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  token,
		"refresh_token": nil,
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
