package routes

import (
	"github.com/gin-gonic/gin"
	"crypto/rsa"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
)

var rsaCertificate *rsa.PrivateKey

func fatal(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ConfigureRoutes(r *gin.Engine) {
	r.POST("/token", token)
	r.GET("/health", health)
}

func LoadCertificate() {
	b, err := ioutil.ReadFile("./keys/jwtRS256.key")
	fatal(err)
	rsaCertificate, err = jwt.ParseRSAPrivateKeyFromPEM(b)
	fatal(err)
}
