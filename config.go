package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	username string
	password string
	company  string
	url      string
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
		url:      os.Getenv("URL"),
	}

	return nil
}

func GetConfig() Config {

	return config
}
