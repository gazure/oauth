package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gazure/oauth/token-generators"
	"github.com/gazure/oauth/models"
	"github.com/satori/go.uuid"
)

const httpAccept = "Accept"

func showLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "login",
		},
	)
}

func showRegistrationPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"title": "Register",
		},
	)
}

func performLogin(c *gin.Context) {
	showLoginPage(c)
}

func performRegistration(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if user := models.CreateUser(username, password); user != nil{
		id, _ := uuid.FromBytes(user.Id)
		token, _ := token_generators.IssueJwt(id.String(), rsaCertificate)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, http.StatusOK, gin.H{
			"title": "successful registration",
			"status": "SUCCESS",
		}, "login-successful.html")
	} else {
		data := gin.H{
			"status":       "Registration failed",
			"ErrorTitle":   "Registration failed",
			"ErrorMessage": "We don't really know!",
		}
		render(c, http.StatusBadRequest, data, "register.html")
	}
}

func logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func render(c *gin.Context, statusCode int, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get(httpAccept) {
	case gin.MIMEJSON:
		c.JSON(statusCode, data)
	default:
		c.HTML(statusCode, templateName, data)
	}
}
