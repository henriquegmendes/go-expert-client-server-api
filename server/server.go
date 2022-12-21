package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-expert-client-server-api/cfg"
	"github.com/henriquegmendes/go-expert-client-server-api/server/client"
	"github.com/henriquegmendes/go-expert-client-server-api/server/database"
	"github.com/henriquegmendes/go-expert-client-server-api/server/handler"
	"github.com/henriquegmendes/go-expert-client-server-api/server/repository"
	"github.com/henriquegmendes/go-expert-client-server-api/server/service"
	"log"
	"net/http"
)

func InitServer() {
	ginEngine := gin.Default()

	db := database.ConnectToDatabase()
	exchangeClient := client.NewExchangeClient(http.Client{}, cfg.Env().ExchangeApiBaseURL, cfg.Env().ExchangeApiResponseTimeout)
	exchangeRepository := repository.NewExchangeRepository(db, cfg.Env().DbQueryResponseTimeout)
	exchangeService := service.NewExchangeService(exchangeClient, exchangeRepository)

	handler.LoadExchangeRoutes(ginEngine, exchangeService)

	port := fmt.Sprintf(":%v", cfg.Env().Port)
	err := ginEngine.Run(port)
	if err != nil {
		log.Fatalf("error to init server at PORT %s: %s", port, err.Error())
	}
}
