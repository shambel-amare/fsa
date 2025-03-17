package util

import (
	"fsa/internal/dto"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {

	ctx.JSON(
		statusCode,
		struct {
			Data interface{} `json:"data"`
		}{
			Data: data,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *dto.ErrorResponse) {
	ctx.AbortWithStatusJSON(err.Code, dto.Response{
		OK:    false,
		Error: err,
	})
}
