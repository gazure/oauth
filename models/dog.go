package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Dog struct {
	Id        []byte
	Name      string
	Breed     string
	OwnerId   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Dog) ToDTO() gin.H{
	return gin.H{
		"Name": d.Name,
		"Breed": d.Breed,
		"AddedOn": d.CreatedAt.Format(time.RFC822),
	}
}

func CreateDog(name, breed string, owner *User) *Dog {
	var dog *Dog
	dog = &Dog{
		Id:        newUuid(),
		Name:      name,
		Breed:     breed,
		OwnerId:   owner.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(dog)
	return dog
}

func GetDogs(owner *User) []Dog {
	dogs := make([]Dog, 0)
	db.Where(&Dog{OwnerId: owner.Id}).Find(&dogs)
	return dogs
}
