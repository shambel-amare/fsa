package cmd

import "github.com/gin-gonic/gin"

func InitRoutes(group *gin.RouterGroup, handler Handler) {
	group.POST("/flights/search", handler.Search.SearchFlight)

}
