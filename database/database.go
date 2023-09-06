package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	DBName    string
	Charset   string
	ParseTime string
	Loc       string
}

func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := DBConfig{
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		Charset:   os.Getenv("DB_CHARSET"),
		ParseTime: os.Getenv("DB_PARSETIME"),
		Loc:       os.Getenv("DB_LOC"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", config.Username, config.Password, config.Host, config.Port, config.DBName, config.Charset, config.ParseTime, config.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
	return db
}
