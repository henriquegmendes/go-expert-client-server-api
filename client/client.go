package client

import (
	"context"
	"github.com/henriquegmendes/go-expert-client-server-api/cfg"
	"log"
	"net/http"
)

func InitClient() {
	ctx := context.Background()
	client := NewInternalClient(http.Client{}, cfg.Env().InternalServerBaseURL, cfg.Env().InternalServerResponseTimeout)

	err := DoUSDBRLBidRequestAndSaveResult(ctx, client)
	if err != nil {
		log.Fatalf("error getting and saving bid request: %s", err.Error())
	}
}
