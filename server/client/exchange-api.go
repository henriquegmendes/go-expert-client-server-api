package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/henriquegmendes/go-expert-client-server-api/errors"
	"github.com/henriquegmendes/go-expert-client-server-api/helpers"
	"io"
	"net/http"
	"time"
)

type GetUSDBRExchangeQuoteResponse struct {
	USDBRL USDBRExchangeQuoteResponse `json:"USDBRL"`
}

type USDBRExchangeQuoteResponse struct {
	Code       string `json:"code"`
	CodeIN     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ExchangeClient interface {
	GetUSDBRLExchangeQuote(ctx context.Context) (*GetUSDBRExchangeQuoteResponse, error)
}

type exchangeClient struct {
	client          http.Client
	baseURL         string
	responseTimeout time.Duration
}

func (c *exchangeClient) GetUSDBRLExchangeQuote(ctx context.Context) (*GetUSDBRExchangeQuoteResponse, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.responseTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctxWithTimeout, http.MethodGet, fmt.Sprintf("%s%s", c.baseURL, "/json/last/USD-BRL"), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		if helpers.CheckContextTimedOutError(err) {
			return nil, errors.RequestTimedOutError
		}

		return nil, err
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if helpers.CheckClientWithErrorStatusCode(response.StatusCode) {
		return nil, helpers.ParseClientErrorResponse(response.StatusCode, responseBytes)
	}

	var result GetUSDBRExchangeQuoteResponse
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewExchangeClient(client http.Client, baseURL string, responseTimeout time.Duration) ExchangeClient {
	return &exchangeClient{
		client:          client,
		baseURL:         baseURL,
		responseTimeout: responseTimeout,
	}
}
