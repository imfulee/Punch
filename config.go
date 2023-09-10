package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	username string
	password string
	company  string
}

var config Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	config = Config{
		username: os.Getenv("USERNAME"),
		password: os.Getenv("PASSWORD"),
		company:  os.Getenv("COMPANY"),
	}

	return nil
}

func GetConfig() Config {
	return config
}
