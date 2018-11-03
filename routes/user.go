package routes

import (
	"github.com/gazure/oauth/models"
	"github.com/gazure/oauth/token-generators"
	"github.com/gin-gonic/gin"
	"net/http"
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

func setLoginCookie(c *gin.Context, user *models.User) error {
	token, err := token_generators.IssueJwt(user.GetId(), rsaCertificate)
	if err != nil {
		return err
	}
	c.SetCookie("id", user.GetId(), 3600, "","", false, true)
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)
	return nil
}

func performLogin(c *gin.Context) {
	errorParams := gin.H{
		"status":       "login unsuccessful",
		"ErrorTitle:":  "Login Failed",
		"ErrorMessage": "User name or password does not match our records",
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.GetUser(username)
	if user == nil || !user.PasswordMatch(password) {
		Render(c, 401, "login.html", errorParams)
	} else {
		setLoginCookie(c, user)
		Render(c, 200, "login-successful.html", gin.H{"title": "Login successful"})
	}
}

func performRegistration(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if user := models.CreateUser(username, password); user != nil {
		setLoginCookie(c, user)
		Render(c, http.StatusOK, "login-successful.html", gin.H{
			"title":  "successful registration",
			"status": "SUCCESS",
		})
	} else {
		data := gin.H{
			"status":       "Registration failed",
			"ErrorTitle":   "Registration failed",
			"ErrorMessage": "We don't really know!",
		}
		Render(c, http.StatusBadRequest, "register.html", data)
	}
}

func logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
