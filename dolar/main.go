package main

import (
	"context"
	"log"
	"os"

	"github.com/laurati/dolar_exchange_rate/internal/configuration"
	"github.com/laurati/dolar_exchange_rate/internal/database"
	"github.com/laurati/dolar_exchange_rate/internal/handler"
	"github.com/laurati/dolar_exchange_rate/internal/repository"
	"github.com/laurati/dolar_exchange_rate/internal/router"
	"github.com/laurati/dolar_exchange_rate/internal/server"
	"gorm.io/driver/sqlite"
)

func main() {

	configuration.EnvironmentSetup()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Server Port: %v", os.Getenv("PORT_DOLAR"))
	log.Println("Starting process API...")

	sqliteDialect := sqlite.Open("./dolar/sqlite/sqlite.db")
	db := database.ConnectDatabase(ctx, sqliteDialect)
	database.Migrate(db)

	repo := repository.NewExchangeRateRepo(db)
	handler := handler.NewExchangeHandler(repo)
	router := router.InitializeRouter(handler)

	server := server.NewServer(":"+os.Getenv("PORT_DOLAR"), router)
	server.Start()

}
