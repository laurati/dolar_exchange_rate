package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/laurati/exchange_rate/internal/entity"
	"github.com/laurati/exchange_rate/internal/repository"
)

type ExchangeHandler struct {
	Repo *repository.ExchangeRateRepo
}

func NewExchangeHandler(Repo *repository.ExchangeRateRepo) *ExchangeHandler {
	return &ExchangeHandler{
		Repo: Repo,
	}
}

func (h *ExchangeHandler) GetExchangeRateByCodeAndSave(c *gin.Context) {

	code := c.Param("code")

	exchangeUrl := os.Getenv("EXCHANGE_API_URL")
	exchangeRateUrl := exchangeUrl + code

	request, err := http.NewRequest("GET", exchangeRateUrl, nil)
	if err != nil {
		log.Println(err.Error())
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("Server error: %s\n", response.Status)
	}

	var exchangeRateMap map[string]entity.ExchangeRate
	err = json.Unmarshal(body, &exchangeRateMap)
	if err != nil {
		log.Println("error unmarshal:", err)
		return
	}

	var key string
	for chave := range exchangeRateMap {
		key = chave
		break
	}

	exchangeRate := exchangeRateMap[key]

	h.Repo.SaveExchangeRateRepo(c.Request.Context(), &exchangeRate)

	exResponse := fmt.Sprintf("The current %v exchange rate is %v", exchangeRate.Code, exchangeRate.Bid)

	c.JSON(http.StatusOK, exResponse)

}

func (h *ExchangeHandler) GetAllExchangeRate(c *gin.Context) {
	exchanges, err := h.Repo.ReadExchangeRateRepo(c.Request.Context())

	if err == errors.New("list exchange rates is empty") {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, exchanges)

}

func (h *ExchangeHandler) GetDolarExchangeRate(c *gin.Context) {

	dolarUrl := os.Getenv("EXCHANGE_DOLAR_API_URL")

	request, err := http.NewRequest("GET", dolarUrl, nil)
	if err != nil {
		log.Println(err.Error())
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("Server error: %s\n", response.Status)
		c.JSON(response.StatusCode, nil)
	}

	var exchangeRateMap map[string]entity.ExchangeRate
	err = json.Unmarshal(body, &exchangeRateMap)
	if err != nil {
		log.Println("error unmarshal:", err)
		return
	}

	var key string
	for chave := range exchangeRateMap {
		key = chave
		break
	}

	exchangeRate := exchangeRateMap[key]

	c.JSON(http.StatusOK, exchangeRate)

}
