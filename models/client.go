package models

import (
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"crypto/rand"
	"encoding/hex"
)

type Client struct {
	Id []byte
	Name string
	Description string
	Secret []byte
}

func (c *Client) GetId() string {
	id, _ := uuid.FromBytes(c.Id)
	return id.String()
}

func (c *Client) GetSecret() string {
	return hex.EncodeToString(c.Secret)
}

func (c *Client) SecretMatches(secret string) bool {
	return c.GetSecret() == secret
}

func (c *Client) ToDTO() gin.H {
	return gin.H{
		"id": c.GetId(),
		"secret": c.GetSecret(),
		"name": c.Name,
		"description": c.Description,
	}
}

func CreateClient(name string, description string) *Client {
	id, _ := uuid.NewV4().MarshalBinary()
	secret := make([]byte, 32)
	rand.Read(secret)
	client := &Client{Id: id, Name: name, Description: description, Secret: secret}
	db.Create(client)
	return client
}

func GetClient(id uuid.UUID) *Client {
	var client Client
	db.Where(&Client{Id: id.Bytes()}).First(&client)
	return &client
}