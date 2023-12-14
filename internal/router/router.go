package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurati/exchange_rate/internal/handler"
)

func InitializeRouter(exchange *handler.ExchangeHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	exchangeRate := router.Group("/exchange")
	{
		exchangeRate.GET("/:code", exchange.GetExchangeRateByCodeAndSave)
		exchangeRate.GET("/", exchange.GetAllExchangeRate)
		exchangeRate.GET("/dolar/", exchange.GetDolarExchangeRate)
	}

	return router
}
