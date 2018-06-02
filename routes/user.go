package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gazure/oauth/token-generators"
	"github.com/gazure/oauth/models"
	"github.com/satori/go.uuid"
)

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

	if user, err := registerNewUser(username, password); err == nil {
		id, _ := uuid.FromBytes(user.Id)
		id.String()
		token, _ := token_generators.IssueJwt(id.String(), rsaCertificate)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.HTML(
			http.StatusOK,
			"login-successful.html",
			gin.H{
				"title": "Successful registration",
			},
		)
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle": "Registration Failed",
			"ErrorMessage": err.Error(),
		})
	}
}

func registerNewUser(username string, password string) (user models.User, err error) {
	err = nil
	user = models.CreateUser(username, password)
	return user, err
}

func logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

