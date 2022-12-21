package client

import (
	"context"
	"log"
)

func DoUSDBRLBidRequestAndSaveResult(ctx context.Context, client InternalClient) error {
	result, err := client.GetInternalUSDBRLExchangeQuote(ctx)
	if err != nil {
		return err
	}

	err = SaveLatestQuoteInFile(result.Bid)
	if err != nil {
		return err
	}

	log.Printf("new quote saved succesfully in ./result/cotacao.txt")

	return nil
}
