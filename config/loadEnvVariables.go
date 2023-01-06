package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading environment variables, %v", err)
	}
}