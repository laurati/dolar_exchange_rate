package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/laurati/dolar_exchange_rate/internal/repository"
)

type ExchangeHandler struct {
	Repo *repository.ExchangeRateRepo
}

func NewExchangeHandler(Repo *repository.ExchangeRateRepo) *ExchangeHandler {
	return &ExchangeHandler{
		Repo: Repo,
	}
}

func (h *ExchangeHandler) GetExchangeRateByCode(c *gin.Context) {

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
		log.Printf("Erro no servidor: %s\n", response.Status)
	}

	laura := make(map[string]interface{})

	err = json.Unmarshal(body, &laura)
	if err != nil {
		fmt.Println("Erro ao fazer unmarshal:", err)
		return
	}

	fmt.Println(laura)

	var key string
	for chave := range laura {
		key = chave
		break
	}

	fmt.Println(key)

	if subMap, ok := laura[key].(map[string]interface{}); ok {
		if bid, ok := subMap["bid"].(string); ok {
			fmt.Println("O valor de bid é:", bid)
		} else {
			fmt.Println("A chave 'bid' não foi encontrada no mapa interno.")
		}
	} else {
		fmt.Println("A chave ", key, " não foi encontrada no mapa principal.")
	}

	// TODO:salvar no banco

	c.JSON(http.StatusOK, laura)

}
