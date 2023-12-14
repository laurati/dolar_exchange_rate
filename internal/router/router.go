package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurati/dolar_exchange_rate/internal/handler"
)

func InitializeRouter(exchange *handler.ExchangeHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	exchangeRate := router.Group("/exchange")
	{
		exchangeRate.GET("/:code", exchange.GetExchangeRateByCode)
		exchangeRate.GET("/", exchange.GetAllExchangeRate)
	}

	return router
}
