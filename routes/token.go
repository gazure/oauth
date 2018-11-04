package routes

import (
	"errors"
	"github.com/gazure/oauth/models"
	"github.com/gazure/oauth/token-generators"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

const paramClientId = "client_id"
const paramClientSecret = "client_secret"
const postParamScope = "scope"
const paramGrantType = "grant_type"
const paramUsername = "username"
const paramPassword = "password"

const grantTypeClientCredentials = "client_credentials"
const grantTypePassword = "password"

var handlerMap = map[string]gin.HandlerFunc{
	grantTypeClientCredentials: handleClientCredentials,
	grantTypePassword:          handlePassword,
}

type oauthEntity interface {
	GetId() string
}

func requestError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}

func badRequest(c *gin.Context, message string) {
	requestError(c, http.StatusBadRequest, message)
}

func unauthorizedRequest(c *gin.Context) {
	requestError(c, http.StatusUnauthorized, "Not authorized.")
}

func renderTokenResponse(c *gin.Context, sub oauthEntity) {
	token, err := token_generators.IssueJwt(sub.GetId(), rsaCertificate)
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

func validateClientCredentials(clientId string, clientSecret string) (client *models.Client, err error) {
	clientIdUUID, err := uuid.FromString(clientId)
	if err != nil {
		return
	}
	client = models.GetClient(clientIdUUID)
	if client.SecretMatches(clientSecret) {
		return
	}
	err = errors.New("invalid client id or secret")
	return
}

func handleClientCredentials(c *gin.Context) {
	clientId := c.PostForm(paramClientId)
	clientSecret := c.PostForm(paramClientSecret)
	user, err := validateClientCredentials(clientId, clientSecret)
	if err != nil {
		log.Println(err)
		unauthorizedRequest(c)
		return
	}
	renderTokenResponse(c, user)
}

func handlePassword(c *gin.Context) {
	username := c.PostForm(paramUsername)
	password := c.PostForm(paramPassword)

	user := models.GetUser(username)
	if user == nil || !user.PasswordMatch(password) {
		unauthorizedRequest(c)
		return
	}
	renderTokenResponse(c, user)
}

func token(c *gin.Context) {
	grantType := c.PostForm(paramGrantType)
	if handlerFunction, ok := handlerMap[grantType]; ok {
		handlerFunction(c)
	} else {
		badRequest(c, "invalid grant type specified")
	}
}
