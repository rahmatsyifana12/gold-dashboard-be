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

type GetUserGoldAssetsResponse struct {
	GoldAssets []GetUserGoldAssetResponse `json:"gold_assets"`
}

type GetUserGoldAssetResponse struct {
	ID            uint    `json:"id"`
	UserID        uint    `json:"user_id"`
	Brand         string  `json:"brand"`
	UnitGram      float64 `json:"unit_gram"`
	CertificateNo string  `json:"certificate_no"`
	Status        string  `json:"status"`
	BoughtPrice   float64 `json:"bought_price"`
	SoldPrice     float64 `json:"sold_price"`
	BuyDate       string  `json:"buy_date"`
	SellDate      string  `json:"sell_date"`
}
