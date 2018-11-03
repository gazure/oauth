package routes

import (
	"github.com/gazure/oauth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func showDogs(c *gin.Context) {
	ownerId, err := c.Cookie("id")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/u/login")
		return
	}
	owner := models.GetUserById(ownerId)
	dogs := models.GetDogs(owner)
	dogsDTOs := make([]gin.H, len(dogs))
	for i, dog := range dogs {
		dogsDTOs[i] = dog.ToDTO()
	}

	params := gin.H{
		"dogs": dogsDTOs,
	}
	Render(c, 200, "dogs.html", params)
}

func showCreateDogForm(c *gin.Context) {
	Render(c, 200, "add-dog.html", gin.H{})
}

func handleCreateDog(c *gin.Context) {
	dogName := c.PostForm("name")
	dogBreed := c.PostForm("breed")
	ownerId, _ := c.Cookie("id")
	owner := models.GetUserById(ownerId)

	if owner == nil {
		Render(c, http.StatusBadRequest, "add-dog.html", gin.H{
			"ErrorTitle": "No owner found!",
			"ErrorMessage": "We couldn't find an owner for this dog",
		})
		return
	}

	models.CreateDog(dogName, dogBreed, owner)
	Render(c, 201, "add-dog.html", gin.H{})
}
