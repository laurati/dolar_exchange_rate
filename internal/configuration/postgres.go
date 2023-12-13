package configuration

import (
	"fmt"
	"log"
	"os"
)

func GetPostgresConnectionString() string {
	log.Printf("Postgres port: %v", os.Getenv("DB_POSTGRES_PORT"))
	log.Printf("Postgres user: %v", os.Getenv("DB_POSTGRES_USER"))
	log.Printf("Postgres database: %v", os.Getenv("DB_POSTGRES_NAME"))

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_NAME"))
}
