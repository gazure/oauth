package routes

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/gazure/oauth/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var rsaCertificate *rsa.PrivateKey

func fatal(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func GetTemplateParameters(c *gin.Context) gin.H {
	return gin.H{
		"isLoggedIn": middleware.IsLoggedIn(c),
	}
}

func MergeTemplateParameters(c *gin.Context, overrides gin.H) gin.H {
	defaults := GetTemplateParameters(c)
	for k, v := range overrides {
		defaults[k] = v
	}
	return defaults
}

func Render(c *gin.Context, code int, name string, params gin.H) {
	mergedParams := MergeTemplateParameters(c, params)

	switch c.Request.Header.Get(httpAccept) {
	case gin.MIMEJSON:
		c.JSON(code, mergedParams)
	default:
		c.HTML(code, name, mergedParams)
	}
}

func index(c *gin.Context) {
	params := gin.H{
		"title": "Home Page",
	}
	Render(c, 200, "index.html", params)
}

func ConfigureRoutes(r *gin.Engine) {
	r.Use(middleware.SetUserStatus)
	r.GET("/", index)
	r.GET("/health", health)
	oauthRoutes := r.Group("/oauth")
	{
		oauthRoutes.POST("/token", token)
		oauthRoutes.POST("/client/register", registerNewClient)
		oauthRoutes.GET("/clients", listClients)
		oauthRoutes.POST("/clients/newSecret", generateNewSecret)
	}
	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/logout", middleware.EnsureLoggedIn, logout)
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn, showLoginPage)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn, showRegistrationPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn, performLogin)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn, performRegistration)
	}
	dogRoutes := r.Group("/dog")
	{
		dogRoutes.GET("/show", middleware.EnsureLoggedIn, showDogs)
		dogRoutes.GET("/create", middleware.EnsureLoggedIn, showCreateDogForm)
		dogRoutes.POST("/create", middleware.EnsureLoggedIn, handleCreateDog)
	}
}

func LoadCertificate() {
	b, err := ioutil.ReadFile("./keys/jwtRS256.key")
	fatal(err)
	rsaCertificate, err = jwt.ParseRSAPrivateKeyFromPEM(b)
	fatal(err)
}
