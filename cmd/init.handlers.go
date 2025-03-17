package cmd

import (
	"fsa/internal/domain"
	"fsa/internal/handlers/rest"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// this way of init lets ass add more api handlers at one place
type Handler struct {
	Search *rest.SearchHandler
}

func InitHandlers(svc domain.SearchService, log zap.Logger) Handler {
	sh := rest.NewSearchHandler(svc, viper.GetDuration("server.read_timeout"), log)
	return Handler{
		Search: sh,
	}
}
