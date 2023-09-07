// database/database.go

package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	DBName    string
	Charset   string
	ParseTime bool
	Loc       string
}

var DB *gorm.DB

func InitDatabase(config DbConfig) *gorm.DB {
	parseTimeStr := "False"
	if config.ParseTime {
		parseTimeStr = "True"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", config.Username, config.Password, config.Host, config.Port, config.DBName, config.Charset, parseTimeStr, config.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
	return db
}
