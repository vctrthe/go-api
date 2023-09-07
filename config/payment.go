package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".payment_env")
	if err != nil {
		log.Fatal("Error loading .payment_env file")
	}
}

var (
	MidtransServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	MidtransClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
)
