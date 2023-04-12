package http

import (
	"net/http"

	"github.com/gas_price/domain"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// GasPriceHandler  represent the httphandler for GasPrice
type GasPriceHandler struct {
	GPUsecase domain.GasPriceUsecase
}

func NewGasPriceHandler(e *echo.Echo, us domain.GasPriceUsecase) {
	handler := &GasPriceHandler{
		GPUsecase: us,
	}
	e.GET("/gas_price", handler.FetchGasPrice)
}

func (gh *GasPriceHandler) FetchGasPrice(c echo.Context) error {
	ctx := c.Request().Context()

	listGP, err := gh.GPUsecase.FetchGasPriceList(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listGP)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
