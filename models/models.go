package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"github.com/satori/go.uuid"
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

func Migrate() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Client{})
}

func newUuid() []byte {
	id, _ := uuid.NewV4().MarshalBinary()
	return id
}