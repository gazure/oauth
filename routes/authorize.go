package routes

import (
	"github.com/gazure/oauth/models"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"net/url"
)

func showAuthorize(c *gin.Context) {
	clientId := c.Query("client_id")
	scope := c.Query("scope")
	state := c.Query("state")
	redirectUrl := c.Query("redirect_url")

	clientUuid, err := uuid.FromString(clientId)
	if err != nil {
		Render(c, 400, "index.html", gin.H{})
		return
	}

	client := models.GetClient(clientUuid)
	if client == nil {
		Render(c, 400, "index.html", gin.H{})
		return
	}

	if redirectUrl == "" {
		redirectUrl = client.RedirectURI
	}

	Render(c, 200, "authorize.html", gin.H{
		"clientId": clientId,
		"scope": scope,
		"state": state,
		"redirectUrl": redirectUrl,
	})
}

func handleAuthorize(c *gin.Context) {
	redirectUrl := c.PostForm("redirectUrl")
	clientId := c.PostForm("clientId")
	code := clientId + ":code"
	state := c.PostForm("state")

	queryString := "?code=" + url.QueryEscape(code) + "&state=" + url.QueryEscape(state)
	c.Redirect(http.StatusFound, redirectUrl + queryString)
}
