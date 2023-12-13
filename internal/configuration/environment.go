package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvironmentSetup() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading env file %v", err)
		return
	}
	log.Println("YOUR ENV FILE IS LOADED")
}
