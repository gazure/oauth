package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func Init() error {
	var err error
	for i := 10; i > 0; i-- {
		db, err = gorm.Open("mysql", "root:test123@tcp(mysql:3306)/oauth?charset=utf8&parseTime=True&loc=Local")
		if err == nil {
			return nil
		}
		log.Println(err)
		time.Sleep(time.Second * 3)
	}
	return err
}

type User struct {
	Id       []byte
	Name     string
	PasswordHash string
}

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func Migrate() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

func CreateUser(name string, password string) User {
	var user User
	if user = GetUser(name); len(user.Id) != 0 {
		return user
	}
	id, _ := uuid.NewV4().MarshalBinary()
	passwordHash, _ := generatePasswordHash(password)
	user = User{Id: id, Name: name, PasswordHash: passwordHash}
	db.Create(user)
	return user
}

func GetUser(name string) User {
	var user User
	db.Where(&User{Name: name}).First(&user)
	return user
}

func (u *User) GetId() string {
	id, _ := uuid.FromBytes(u.Id)
	return id.String()
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
