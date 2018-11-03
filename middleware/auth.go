package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const isLoggedIn = "is_logged_in"

func SetUserStatus(c *gin.Context) {
	if token, err := c.Cookie("token"); err == nil || token != "" {
		c.Set(isLoggedIn, true)
	} else {
		c.Set(isLoggedIn, false)
	}
}

func IsLoggedIn(c *gin.Context) bool {
	loggedInInteface, _ := c.Get(isLoggedIn)
	return loggedInInteface.(bool)
}

func EnsureLoggedIn(c *gin.Context) {
	if !IsLoggedIn(c) {
		c.Redirect(http.StatusFound, "/u/login")
	}
}

func EnsureNotLoggedIn(c *gin.Context) {
	if IsLoggedIn(c) {
		c.Redirect(http.StatusFound, "/")
	}
}
