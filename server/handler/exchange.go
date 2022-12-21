package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-expert-client-server-api/server/service"
)

type exchangeHandler struct {
	router  *gin.Engine
	service service.ExchangeService
}

func LoadExchangeRoutes(router *gin.Engine, exchangeService service.ExchangeService) {
	handler := exchangeHandler{
		router:  router,
		service: exchangeService,
	}

	router.POST("/cotacao", handler.GetAndSaveBid)
	router.GET("/cotacao", handler.GetAll)
}

func (h *exchangeHandler) GetAndSaveBid(ctx *gin.Context) {
	response, err := h.service.GetAndSaveBid(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "unexpected error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"bid": response.Bid,
	})
}

func (h *exchangeHandler) GetAll(ctx *gin.Context) {
	response, err := h.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "unexpected error",
		})
		return
	}

	ctx.JSON(200, response)
}
