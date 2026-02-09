package handlers

import (
	"gold-dashboard-be/internal/constants"
	"gold-dashboard-be/internal/dtos"
	"gold-dashboard-be/internal/pkg/helpers"
	"gold-dashboard-be/internal/pkg/responses"
	"gold-dashboard-be/internal/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type GoldAssetHandler interface {
	CreateGoldAsset(c echo.Context) (err error)
}

type GoldAssetHandlerImpl struct {
	usecase *usecases.Usecase
}

func NewGoldAssetHandler(ioc di.Container) GoldAssetHandler {
	return &GoldAssetHandlerImpl{
		usecase: ioc.Get(constants.Usecase).(*usecases.Usecase),
	}
}

func (t *GoldAssetHandlerImpl) CreateGoldAsset(c echo.Context) (err error) {
	var (
		ctx    = c.Request().Context()
		params dtos.CreateGoldAssetRequest
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	if err = c.Validate(&params); err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Validation error").
			SendErrorResponse(c)
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims").
			SendErrorResponse(c)
	}

	err = t.usecase.GoldAsset.CreateGoldAsset(ctx, claims, params)

	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		WithMessage("Successfully created a new gold asset").
		Send(c)
}
