package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ExchangeRate struct {
	Code       string     `json:"code"`
	Codein     string     `json:"codein"`
	Name       string     `json:"name"`
	High       string     `json:"high"`
	Low        string     `json:"low"`
	VarBid     string     `json:"varBid"`
	PctChange  string     `json:"pctChange"`
	Bid        string     `json:"bid"`
	Ask        string     `json:"ask"`
	Timestamp  string     `json:"timestamp"`
	CreateDate CustomTime `json:"create_date"`
}

func (ExchangeRate) TableName() string {
	return "exchange_rate"
}

type CustomTime struct {
	time.Time
}

// Scan implements the interface sql.Scanner.
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	if v, ok := value.(time.Time); ok {
		ct.Time = v
		return nil
	}
	return fmt.Errorf("cannot convert %v to CustomTime", value)
}

// Value implements the interface driver.Valuer.
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	const customTimeLayout = "2006-01-02 15:04:05"
	s := string(b)
	t, err := time.Parse(`"`+customTimeLayout+`"`, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
