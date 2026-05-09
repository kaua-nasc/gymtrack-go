package config

import "github.com/joho/godotenv"

func LoadEnvironmentVariable() {
	_ = godotenv.Load()
}
