package main

import (
	"context"
	"log"
	"os"

	"github.com/laurati/exchange_rate/internal/configuration"
	"github.com/laurati/exchange_rate/internal/database"
	"github.com/laurati/exchange_rate/internal/handler"
	"github.com/laurati/exchange_rate/internal/repository"
	"github.com/laurati/exchange_rate/internal/router"
	"github.com/laurati/exchange_rate/internal/server"
	"gorm.io/driver/postgres"
)

func main() {

	configuration.EnvironmentSetup()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Server Port: %v", os.Getenv("PORT"))
	log.Println("Starting process API...")

	postgresDsn := configuration.GetPostgresConnectionString()
	dbPostgres := database.ConnectDatabase(ctx, postgres.Open(postgresDsn))
	database.Migrate(dbPostgres)

	repo := repository.NewExchangeRateRepo(dbPostgres)
	handler := handler.NewExchangeHandler(repo)
	router := router.InitializeRouter(handler)

	server := server.NewServer(":"+os.Getenv("PORT"), router)
	server.Start()

}
