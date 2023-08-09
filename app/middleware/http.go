package middleware

import (
	"github.com/dstgo/wilson/app/pkg/httputil"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func UseCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// NotFoundHandler
// return gin.HandlerFunc
func NotFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httputil.Failed(ctx, 4040, httputil.NotFoundError(errors.New("requested path not found")))
	}
}

func NoMethodHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httputil.Failed(ctx, 4050, httputil.NotFoundError(errors.New("requested http method not supported")))
	}
}
