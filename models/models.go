package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"github.com/satori/go.uuid"
	"os"
)

var db *gorm.DB
const myslqCstringEnvvar = "MYSQL_CSTRING"

func Init() error {
	var err error
	mysqlCString, ok := os.LookupEnv(myslqCstringEnvvar)
	if !ok {
		mysqlCString = "root:test123@tcp(mysql:3306)/oauth?charset=utf8&parseTime=True&loc=Local"
	}
	for i := 10; i > 0; i-- {
		db, err = gorm.Open("mysql", mysqlCString)
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
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Dog{})
}

func newUuid() []byte {
	id, _ := uuid.NewV4().MarshalBinary()
	return id
}
