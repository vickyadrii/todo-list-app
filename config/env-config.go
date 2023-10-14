package config

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load()
	if (err) != nil {
		panic("Can't load environment variables from .env file")
	}
}
