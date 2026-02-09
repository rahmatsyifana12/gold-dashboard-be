package dtos

type CreateGoldAssetRequest struct {
	Brand         string   `json:"brand" validate:"required"`
	UnitGram      float64  `json:"unit_gram" validate:"required"`
	CertificateNo string   `json:"certificate_no" validate:"required"`
	Status        string   `json:"status" validate:"required,oneof=owned sold"`
	BoughtPrice   float64  `json:"bought_price" validate:"required"`
	SoldPrice     *float64 `json:"sold_price"`
	BuyDate       string   `json:"buy_date" validate:"required"`
	SellDate      *string  `json:"sell_date"`
}
