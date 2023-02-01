package main

import "github.com/joho/godotenv"

func LoadEnv() error {
	// Load environment variables
	return godotenv.Load(".env")
}
