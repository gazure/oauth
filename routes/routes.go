package routes

import (
	"github.com/gin-gonic/gin"
	"crypto/rsa"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/gazure/oauth/middleware"
)

var rsaCertificate *rsa.PrivateKey

func fatal(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func ConfigureRoutes(r *gin.Engine) {
	r.Use(middleware.SetUserStatus)
	r.GET("/", index)
	r.GET("/health", health)
	oauthRoutes := r.Group("/oauth")
	{
		oauthRoutes.POST("/token", token)
		oauthRoutes.POST("/client/register", registerNewClient)
	}
	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/logout", middleware.EnsureLoggedIn, logout)
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn, showLoginPage)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn, showRegistrationPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn, performLogin)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn, performRegistration)
	}
}

func LoadCertificate() {
	b, err := ioutil.ReadFile("./keys/jwtRS256.key")
	fatal(err)
	rsaCertificate, err = jwt.ParseRSAPrivateKeyFromPEM(b)
	fatal(err)
}
