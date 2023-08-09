package httputil

import (
	"github.com/gin-gonic/gin"
)

const (
	RequestIdKey = "X-Request-ID"
)

func SetRequestId(ctx *gin.Context, id string) {
	ctx.Set(RequestIdKey, id)
	ctx.Writer.Header().Set(RequestIdKey, id)
}

func GetRequestId(ctx *gin.Context) (requestId string) {
	return ctx.GetString(RequestIdKey)
}
