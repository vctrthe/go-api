// config/db.go

package config

import (
	"log"
	"os"
	"strconv" // Add strconv package for boolean conversion

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	DBName    string
	Charset   string
	ParseTime bool // Changed ParseTime to bool
	Loc       string
}

func init() {
	err := godotenv.Load(".db_env")
	if err != nil {
		log.Fatal("Error loading .db_env file")
	}
}

var (
	DbUsername     = os.Getenv("DB_USERNAME")
	DbPassword     = os.Getenv("DB_PASSWORD")
	DbHost         = os.Getenv("DB_HOST")
	DbPort         = os.Getenv("DB_PORT")
	DbName         = os.Getenv("DB_NAME")
	DbCharset      = os.Getenv("DB_CHARSET")
	DbParseTimeStr = os.Getenv("DB_PARSETIME")
	DbLoc          = os.Getenv("DB_LOC")
)

func ParseBoolEnv(envVar string) bool {
	value, err := strconv.ParseBool(envVar)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", envVar, err)
	}
	return value
}

// Initialize the ParseTime boolean from the environment variable
var DbParseTime = ParseBoolEnv(DbParseTimeStr)
