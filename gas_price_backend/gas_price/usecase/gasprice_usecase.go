package usecase

import (
	"context"
	"time"

	"github.com/gas_price/domain"
)

type gasPriceUsecase struct {
	gpRepo         domain.GasPriceRepository
	contextTimeout time.Duration
}

func NewGasPriceUsecase(gpRepo domain.GasPriceRepository, timeout time.Duration) domain.GasPriceUsecase {
	return &gasPriceUsecase{
		gpRepo:         gpRepo,
		contextTimeout: timeout,
	}
}

func (gp *gasPriceUsecase) FetchGasPriceList(ctx context.Context) ([]*domain.GasPrice, error) {
	ctx, cancel := context.WithTimeout(ctx, gp.contextTimeout)
	defer cancel()
	gasPrices, err := gp.gpRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return gasPrices, nil
}

func (gp *gasPriceUsecase) GetGasPriceById(ctx context.Context, id int64) (*domain.GasPrice, error) {
	ctx, cancel := context.WithTimeout(ctx, gp.contextTimeout)
	defer cancel()
	gasPrice, err := gp.gpRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}
func (g *gasPriceUsecase) AddGasPrice(ctx context.Context, gp *domain.GasPrice) error {
	ctx, cancel := context.WithTimeout(ctx, g.contextTimeout)
	defer cancel()
	err := g.gpRepo.Store(ctx, gp)
	if err != nil {
		return err
	}
	return nil
}
func (g *gasPriceUsecase) UpdateGasPrice(ctx context.Context, gp *domain.GasPrice) error {
	ctx, cancel := context.WithTimeout(ctx, g.contextTimeout)
	defer cancel()
	err := g.gpRepo.Update(ctx, gp)
	if err != nil {
		return err
	}
	return nil
}
func (gp *gasPriceUsecase) DeleteGasPrice(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, gp.contextTimeout)
	defer cancel()
	err := gp.gpRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
