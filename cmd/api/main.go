package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/laurati/dolar_exchange_rate/internal/configuration"
	"github.com/laurati/dolar_exchange_rate/internal/database"
	"github.com/laurati/dolar_exchange_rate/internal/server"
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

	fmt.Println(dbPostgres)

	// repo := repository.NewDetailsRepo(dbPostgres)
	// useCase := usecase.NewDetailsUseCase(repo)
	// handler := handler.NewDetailsHandler(useCase)

	// router := router.InitializeRouter(handler)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })
	server := server.NewServer(":"+os.Getenv("PORT"), router)
	server.Start()

}
