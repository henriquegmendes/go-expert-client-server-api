package repository

import (
	"context"
	"github.com/henriquegmendes/go-expert-client-server-api/errors"
	"github.com/henriquegmendes/go-expert-client-server-api/helpers"
	"github.com/henriquegmendes/go-expert-client-server-api/server/models"
	"gorm.io/gorm"
	"time"
)

type ExchangeRepository interface {
	CreateOne(ctx context.Context, exchange *models.Exchange) error
	GetAll(ctx context.Context) ([]models.Exchange, error)
}

type exchangeRepository struct {
	db           *gorm.DB
	queryTimeout time.Duration
}

func (r *exchangeRepository) CreateOne(ctx context.Context, exchange *models.Exchange) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	tx := r.db.WithContext(ctxWithTimeout).Begin()
	tx.Create(exchange)
	if tx.Error != nil {
		tx.Rollback()

		return tx.Error
	}

	err := tx.Commit().Error
	if err != nil {
		if helpers.CheckContextTimedOutError(ctxWithTimeout.Err()) {
			return errors.RequestTimedOutError
		}

		return err
	}

	return nil
}

func (r *exchangeRepository) GetAll(ctx context.Context) ([]models.Exchange, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	tx := r.db.WithContext(ctxWithTimeout).Begin()

	var exchanges []models.Exchange
	tx.Find(&exchanges)
	if tx.Error != nil {
		tx.Rollback()

		return nil, tx.Error
	}

	err := tx.Commit().Error
	if err != nil {
		if helpers.CheckContextTimedOutError(ctxWithTimeout.Err()) {
			return nil, errors.RequestTimedOutError
		}

		return nil, err
	}

	return exchanges, nil
}

func NewExchangeRepository(db *gorm.DB, queryTimeout time.Duration) ExchangeRepository {
	return &exchangeRepository{
		db:           db,
		queryTimeout: queryTimeout,
	}
}
