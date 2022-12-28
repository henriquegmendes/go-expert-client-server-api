package client

import (
	"context"
	"encoding/json"
	"fmt"
	internalErrors "github.com/henriquegmendes/go-expert-client-server-api/errors"
	"github.com/henriquegmendes/go-expert-client-server-api/helpers"
	"io"
	"net/http"
	"time"
)

type InternalExchangeQuoteResponse struct {
	Bid float64 `json:"bid"`
}

type InternalClient interface {
	GetInternalUSDBRLExchangeQuote(ctx context.Context) (*InternalExchangeQuoteResponse, error)
}

type internalClient struct {
	client          http.Client
	baseURL         string
	responseTimeout time.Duration
}

func (c *internalClient) GetInternalUSDBRLExchangeQuote(ctx context.Context) (*InternalExchangeQuoteResponse, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.responseTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctxWithTimeout, http.MethodPost, fmt.Sprintf("%s%s", c.baseURL, "/cotacao"), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		if helpers.CheckContextTimedOutError(err) {
			return nil, internalErrors.RequestTimedOutError
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

	var result InternalExchangeQuoteResponse
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewInternalClient(client http.Client, baseURL string, responseTimeout time.Duration) InternalClient {
	return &internalClient{
		client:          client,
		baseURL:         baseURL,
		responseTimeout: responseTimeout,
	}
}
