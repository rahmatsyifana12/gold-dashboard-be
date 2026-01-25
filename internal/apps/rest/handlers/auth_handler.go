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

type AuthHandler interface {
	Login(c echo.Context) (err error)
	Logout(c echo.Context) (err error)
}

type AuthHandlerImpl struct {
	usecase *usecases.Usecase
}

func NewAuthHandler(ioc di.Container) AuthHandler {
	return &AuthHandlerImpl{
		usecase: ioc.Get(constants.Usecase).(*usecases.Usecase),
	}
}

func (t *AuthHandlerImpl) Login(c echo.Context) (err error) {
	var (
		ctx    = c.Request().Context()
		params dtos.LoginRequest
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

	res, err := t.usecase.Auth.Login(ctx, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully logged in").
		WithData(res).
		Send(c)
}

func (t *AuthHandlerImpl) Logout(c echo.Context) (err error) {
	var (
		ctx = c.Request().Context()
	)

	authClaims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	err = t.usecase.Auth.Logout(ctx, authClaims)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully logged out").
		Send(c)
}
