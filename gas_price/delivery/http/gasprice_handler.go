package http

import (
	"net/http"
	"strconv"

	"github.com/gas_price/domain"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// GasPriceHandler  represent the httphandler for GasPrice
type GasPriceHandler struct {
	GPUsecase domain.GasPriceUsecase
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func NewGasPriceHandler(e *echo.Echo, us domain.GasPriceUsecase) {
	handler := &GasPriceHandler{
		GPUsecase: us,
	}
	e.GET("/gas_price", handler.FetchGasPrice)
	e.POST("/gas_price", handler.Store)
	e.GET("/gas_price/:id", handler.GetByID)
	e.PUT("/gas_price/:id", handler.Update)
	e.DELETE("/gas_price/:id", handler.Delete)
}

func (gh *GasPriceHandler) Update(c echo.Context) (err error) {
	var gasPrice domain.GasPrice
	err = c.Bind(&gasPrice)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = isRequestValid(&gasPrice); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = gh.GPUsecase.UpdateGasPrice(ctx, &gasPrice)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, gasPrice)
}

func (gh *GasPriceHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = gh.GPUsecase.DeleteGasPrice(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

// GetByID will get gas price by given id
func (gh *GasPriceHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := gh.GPUsecase.GetGasPriceById(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *domain.GasPrice) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (gh *GasPriceHandler) Store(c echo.Context) (err error) {
	var gasPrice domain.GasPrice
	err = c.Bind(&gasPrice)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&gasPrice); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = gh.GPUsecase.AddGasPrice(ctx, &gasPrice)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, gasPrice)
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
