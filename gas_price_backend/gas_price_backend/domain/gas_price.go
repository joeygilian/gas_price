package domain

import "context"

type GasPrice struct {
	ID        int     `json:"id"`
	Litre     float64 `json:"litre"`
	Premium   float64 `json:"premium"`
	Pertalite float64 `json:"pertalite"`
}

type GasPriceUsecase interface {
	FetchGasPriceList(ctx context.Context) ([]*GasPrice, error)
	GetGasPriceById(ctx context.Context, id int64) (*GasPrice, error)
	AddGasPrice(ctx context.Context, gp *GasPrice) error
	UpdateGasPrice(ctx context.Context, gp *GasPrice) error
	DeleteGasPrice(ctx context.Context, id int64) error
}

type GasPriceRepository interface {
	Fetch(ctx context.Context) ([]*GasPrice, error)
	GetByID(ctx context.Context, id int64) (*GasPrice, error)
	Store(ctx context.Context, gp *GasPrice) error
	Update(ctx context.Context, gp *GasPrice) error
	Delete(ctx context.Context, id int64) error
}
