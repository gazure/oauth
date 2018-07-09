package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const IsLoggedIn = "is_logged_in"

func SetUserStatus(c *gin.Context) {
	if token, err := c.Cookie("token"); err == nil || token != "" {
		c.Set(IsLoggedIn, true)
	} else {
		c.Set(IsLoggedIn, false)
	}
}

func isLoggedIn(c *gin.Context) bool {
	loggedInInteface, _ := c.Get(IsLoggedIn)
	return loggedInInteface.(bool)
}

func EnsureLoggedIn(c *gin.Context) {
	if !isLoggedIn(c) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func EnsureNotLoggedIn(c *gin.Context) {
	if isLoggedIn(c) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "currently logged in",
		})
	}
}
