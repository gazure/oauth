package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       []byte
	Name     string `gorm:"unique_index"`
	PasswordHash string
}

func (u *User) GetId() string {
	id, _ := uuid.FromBytes(u.Id)
	return id.String()
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(name string, password string) *User {
	var user *User
	if user = GetUser(name); len(user.Id) != 0 {
		return user
	}
	id, _ := uuid.NewV4().MarshalBinary()
	passwordHash, _ := generatePasswordHash(password)
	user = &User{Id: id, Name: name, PasswordHash: passwordHash}
	db.Create(user)
	return user
}

func GetUser(name string) *User {
	var user User
	db.Where(&User{Name: name}).First(&user)
	return &user
}

