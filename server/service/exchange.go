package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-expert-client-server-api/server/client"
	"github.com/henriquegmendes/go-expert-client-server-api/server/models"
	"github.com/henriquegmendes/go-expert-client-server-api/server/repository"
	"strconv"
	"time"
)

type ExchangeService interface {
	GetAndSaveBid(ctx context.Context) (*models.Exchange, error)
	GetAll(ctx *gin.Context) ([]models.Exchange, error)
}

type exchangeService struct {
	exchangeClient     client.ExchangeClient
	exchangeRepository repository.ExchangeRepository
}

func (s *exchangeService) GetAndSaveBid(ctx context.Context) (*models.Exchange, error) {
	exchangeResult, err := s.exchangeClient.GetUSDBRLExchangeQuote(ctx)
	if err != nil {
		return nil, err
	}

	bidNumber, _ := strconv.ParseFloat(exchangeResult.USDBRL.Bid, 64)
	nowUTC := time.Now().UTC()
	newExchangeModel := &models.Exchange{
		Type:      exchangeResult.USDBRL.Name,
		Bid:       bidNumber,
		CreatedAt: nowUTC,
		UpdatedAt: nowUTC,
	}

	err = s.exchangeRepository.CreateOne(ctx, newExchangeModel)
	if err != nil {
		return nil, err
	}

	return newExchangeModel, nil
}

func (s *exchangeService) GetAll(ctx *gin.Context) ([]models.Exchange, error) {
	return s.exchangeRepository.GetAll(ctx)
}

func NewExchangeService(exchangeClient client.ExchangeClient, exchangeRepository repository.ExchangeRepository) ExchangeService {
	return &exchangeService{
		exchangeClient:     exchangeClient,
		exchangeRepository: exchangeRepository,
	}
}
