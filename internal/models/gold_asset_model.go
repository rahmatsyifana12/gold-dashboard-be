package models

import "time"

type GoldAsset struct {
	BaseModel

	UserID        uint       `gorm:"user_id" json:"user_id"`
	Brand         string     `gorm:"brand" json:"brand"`
	UnitGram      float64    `gorm:"unit_gram" json:"unit_gram"`
	CertificateNo string     `gorm:"certificate_no" json:"certificate_no"`
	Status        string     `gorm:"status" json:"status"`
	BoughtPrice   float64    `gorm:"bought_price" json:"bought_price"`
	SoldPrice     *float64   `gorm:"sold_price" json:"sold_price"`
	BuyDate       time.Time  `gorm:"buy_date" json:"buy_date"`
	SellDate      *time.Time `gorm:"sell_date" json:"sell_date"`
}
