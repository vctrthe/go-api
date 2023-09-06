package database

import (
	"fmt"
	"go-api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName, config.DbCharset, config.DbParseTime, config.DbLoc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
	return db
}
