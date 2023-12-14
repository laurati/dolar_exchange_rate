package entity

import (
	"time"
)

type ExchangeRate struct {
	Code       string    `gorm:"primary_key" json:"code"`
	Codein     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varBid"`
	PctChange  string    `json:"pctChange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func (ExchangeRate) TableName() string {
	return "exchange_rate"
}
