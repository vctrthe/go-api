package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".db_env")
	if err != nil {
		log.Fatal("Error loading .db_env file")
	}
}

var (
	DbUsername  = os.Getenv("DB_USERNAME")
	DbPassword  = os.Getenv("DB_PASSWORD")
	DbHost      = os.Getenv("DB_HOST")
	DbPort      = os.Getenv("DB_PORT")
	DbName      = os.Getenv("DB_NAME")
	DbCharset   = os.Getenv("DB_CHARSET")
	DbParseTime = os.Getenv("DB_PARSETIME")
	DbLoc       = os.Getenv("DB_LOC")
)
