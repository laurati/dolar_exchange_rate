package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/laurati/dolar_exchange_rate/internal/entity"
	"gorm.io/gorm"
)

type ExchangeRateRepo struct {
	DB *gorm.DB
}

func NewExchangeRateRepo(DB *gorm.DB) *ExchangeRateRepo {
	return &ExchangeRateRepo{DB}
}

func (b *ExchangeRateRepo) SaveExchangeRateRepo(ctx context.Context, exchangeRate *entity.ExchangeRate) error {

	if err := b.DB.WithContext(ctx).Create(&exchangeRate).Error; err != nil {
		log.Println("error creating an exchange rate", err.Error())
		return err
	}

	var response map[string]interface{}

	data, err := json.Marshal(exchangeRate)
	if err != nil {
		log.Println("error marshal exchange rate", err.Error())
		return err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Println("error unmarshal exchange rate", err.Error())
		return err
	}

	return nil
}

func (b *ExchangeRateRepo) ReadExchangeRateRepo(ctx context.Context) ([]entity.ExchangeRate, error) {
	exchangeRates := []entity.ExchangeRate{}
	err := b.DB.WithContext(ctx).Find(&exchangeRates).Error
	if err != nil {
		log.Println("error finding exchange rates", err.Error())
		return nil, err
	}

	if len(exchangeRates) == 0 {
		return nil, errors.New("list exchange rates is empty")
	}

	return exchangeRates, nil

}
