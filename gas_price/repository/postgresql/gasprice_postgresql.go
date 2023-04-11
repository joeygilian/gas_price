package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gas_price/domain"
)

type postgresqlGasPriceRepository struct {
	Conn *sql.DB
}

func NewPostgresqlGasPriceRepository(conn *sql.DB) domain.GasPriceRepository {
	return &postgresqlGasPriceRepository{conn}
}

func (r *postgresqlGasPriceRepository) Fetch(ctx context.Context) ([]*domain.GasPrice, error) {
	query := `SELECT id, litre, premium, pertalite FROM gas_price`

	rows, err := r.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gasPrices := []*domain.GasPrice{}

	for rows.Next() {
		gp := &domain.GasPrice{}
		err := rows.Scan(&gp.ID, &gp.Litre, &gp.Premium, &gp.Pertalite)
		if err != nil {
			return nil, err
		}

		gasPrices = append(gasPrices, gp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gasPrices, nil
}

func (r *postgresqlGasPriceRepository) GetByID(ctx context.Context, id int64) (*domain.GasPrice, error) {
	query := `SELECT id, litre, premium, pertalite FROM gas_price  WHERE id = $1`
	var gp domain.GasPrice

	err := r.Conn.QueryRowContext(ctx, query, id).Scan(&gp.ID, &gp.Litre, &gp.Premium, &gp.Pertalite)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("gas price with id %d not found", id)
		}
	}

	return &gp, nil
}

func (r *postgresqlGasPriceRepository) Store(ctx context.Context, gp *domain.GasPrice) error {
	query := "INSERT INTO gas_price(litre, premium, pertalite) VALUES ($1, $2, $3) RETURNING id"
	err := r.Conn.QueryRow(query, gp.Litre, gp.Premium, gp.Pertalite).Scan(&gp.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlGasPriceRepository) Update(ctx context.Context, gp *domain.GasPrice) error {
	_, err := r.Conn.Exec("UPDATE gas_prices SET litre=$1, premium=$2, pertalite=$3 WHERE id=$4", gp.Litre, gp.Premium, gp.Pertalite, gp.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlGasPriceRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM gas_price WHERE id=$1"
	_, err := r.Conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
