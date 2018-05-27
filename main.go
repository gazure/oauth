package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
	"io/ioutil"
	"crypto/rsa"
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

var rsaCertificate *rsa.PrivateKey

func fatal(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func loadCertificate() {
	b, err := ioutil.ReadFile("./keys/jwtRS256.key")
	fatal(err)

	rsaCertificate, err = jwt.ParseRSAPrivateKeyFromPEM(b)
	fatal(err)
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

func issueToken(clientId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": clientId,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString(rsaCertificate)
	return tokenStr, err
}

func badRequest(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "no grant_type specified",
	})
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
	token, err := issueToken(clientId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
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
	loadCertificate()
	r := gin.Default()
	r.POST("/token", token)
	r.GET("/health", health)
	r.Run()
}
