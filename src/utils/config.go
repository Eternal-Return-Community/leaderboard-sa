package utils

import (
	"erbs/src/structs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func config() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}
	return nil
}

func token() string {
	if token := os.Getenv("TOKEN"); token != "" {
		return token
	}
	log.Fatal("Token is missing .env!")
	return ""
}

func key() string {
	if key := os.Getenv("KEY"); key != "" {
		return key
	}
	log.Fatal("Key is missing .env!")
	return ""
}

func Env() structs.Env {
	config()

	return structs.Env{
		Token:  "Bot " + token(),
		Key:    key(),
		Prefix: "!",
	}
}
