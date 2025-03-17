package rest

import (
	"context"
	"fsa/internal/domain"
	"fsa/internal/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SearchHandlerConfig struct {
	Service domain.SearchService
}

type SearchHandler struct {
	service    domain.SearchService
	ctxTimeout time.Duration
	log        zap.Logger
}

func NewSearchHandler(
	svc domain.SearchService,
	timeOut time.Duration,
	logger zap.Logger,
) *SearchHandler {
	return &SearchHandler{
		service:    svc,
		ctxTimeout: timeOut,
		log:        logger,
	}
}
func (h *SearchHandler) SearchFlight(ctx *gin.Context) {
	var err error

	c, cancel := context.WithTimeout(ctx, h.ctxTimeout)
	defer cancel()
	req := dto.SearchTravelDTO{}

	err = ctx.ShouldBind(&req)
	if err != nil {
		h.log.Info("unable to bind request data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	response, err := h.service.SearchFlight(c, req)
	if err != nil {
		h.log.Error("unable to search flight", zap.Error(err))
		_ = ctx.Error(err)

		return
	}
	ctx.JSON(
		http.StatusOK,
		struct {
			Data interface{} `json:"data"`
		}{
			Data: response,
		},
	)
}
