package usecases

import (
	"context"
	"gold-dashboard-be/internal/constants"
	"gold-dashboard-be/internal/dtos"
	"gold-dashboard-be/internal/models"
	"gold-dashboard-be/internal/pkg/responses"
	"gold-dashboard-be/internal/repositories"
	"net/http"
	"time"

	"github.com/sarulabs/di"
)

type GoldAssetUseCase interface {
	CreateGoldAsset(ctx context.Context, claims dtos.AuthClaims, params dtos.CreateGoldAssetRequest) (err error)
	GetUserGoldAssets(ctx context.Context, claims dtos.AuthClaims) (response dtos.GetUserGoldAssetsResponse, err error)
}

type GoldAssetUseCaseImpl struct {
	repository *repositories.Repository
}

func NewGoldAssetUseCase(ioc di.Container) *GoldAssetUseCaseImpl {
	return &GoldAssetUseCaseImpl{
		repository: ioc.Get(constants.Repository).(*repositories.Repository),
	}
}

func (u *GoldAssetUseCaseImpl) CreateGoldAsset(ctx context.Context, claims dtos.AuthClaims, params dtos.CreateGoldAssetRequest) (err error) {
	buyDateTime, err := time.Parse("2006-01-02", params.BuyDate)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnprocessableEntity).
			WithMessage("error while parsing buy date")
		return
	}

	newGoldAsset := models.GoldAsset{
		UserID:        claims.UserID,
		Brand:         params.Brand,
		UnitGram:      params.UnitGram,
		CertificateNo: params.CertificateNo,
		Status:        params.Status,
		BoughtPrice:   params.BoughtPrice,
		BuyDate:       buyDateTime,
	}

	if params.Status == constants.StatusSold {
		if params.SoldPrice != nil {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusUnprocessableEntity).
				WithMessage("sold price is required")
			return
		} else if params.SellDate != nil {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusUnprocessableEntity).
				WithMessage("sell date is required")
			return
		}

		sellDateTime, err := time.Parse("2006-01-02", *params.SellDate)
		if err != nil {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusUnprocessableEntity).
				WithMessage("error while parsing sell date")
			return err
		}

		newGoldAsset.SoldPrice = params.SoldPrice
		newGoldAsset.SellDate = &sellDateTime
	}

	err = u.repository.GoldAsset.CreateGoldAsset(ctx, newGoldAsset)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("error while inserting gold asset into database")
		return
	}

	return
}

func (u *GoldAssetUseCaseImpl) GetUserGoldAssets(ctx context.Context, claims dtos.AuthClaims) (response dtos.GetUserGoldAssetsResponse, err error) {
	goldAssets, err := u.repository.GoldAsset.GetGoldAssetsByUserID(ctx, claims.UserID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("error while fetching gold assets from database")
		return
	}

	for _, goldAsset := range goldAssets {
		goldAssetRes := dtos.GetUserGoldAssetResponse{
			ID:            goldAsset.ID,
			Brand:         goldAsset.Brand,
			UnitGram:      goldAsset.UnitGram,
			CertificateNo: goldAsset.CertificateNo,
			Status:        goldAsset.Status,
			BoughtPrice:   goldAsset.BoughtPrice,
			BuyDate:       goldAsset.BuyDate.Format("2006-01-02"),
		}

		if goldAsset.SoldPrice != nil {
			goldAssetRes.SoldPrice = *goldAsset.SoldPrice
		}
		if goldAsset.SellDate != nil {
			goldAssetRes.SellDate = goldAsset.SellDate.Format("2006-01-02")
		}

		response.GoldAssets = append(response.GoldAssets, goldAssetRes)
	}

	return
}
