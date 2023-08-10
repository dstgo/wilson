package middleware

import (
	"fmt"
	"github.com/dstgo/size"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

func UseLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			request   = ctx.Request
			requestId = uuid.NewString()
			startTime = time.Now()
		)

		// next handlers
		ctx.Next()

		// set response header X-Request-ID
		httpx.SetRequestId(ctx, requestId)

		var (
			fullpath     = ctx.FullPath()
			method       = request.Method
			costTime     = time.Now().Sub(startTime).Milliseconds()
			status       = ctx.Writer.Status()
			responseSize = ctx.Writer.Size()
			requestSize  = request.ContentLength

			clientIp = ctx.ClientIP()
			err      = ctx.Err()
		)

		entry := logger.WithContext(ctx).
			WithField(types.LogIpKey, clientIp).
			WithField(types.LogHttpMethodKey, method).
			WithField(types.LogRequestPathKey, fullpath).
			WithField(types.LogHttpStatusKey, status).
			WithField(types.LogRequestIdKey, requestId).
			WithField(types.LogRequestCostKey, fmt.Sprintf("%d ms", costTime)).
			WithField(types.LogHttpContentLength, closeSize(requestSize)).
			WithField(types.LogHttpResponseLength, closeSize(int64(responseSize)))

		// log by status
		switch {
		case err == nil && 100 < status && status < 400:
			entry.Infoln()
		case status >= 400:
			entry.WithError(err)
			entry.Errorln()
		}
	}
}

func closeSize(s int64) string {
	meta := size.NewSize(float64(s), size.B)
	data := meta.Data
	switch {
	case data >= float64(size.MB):
		return size.ParseTargetSize(meta.String(), size.MB).String()
	case data >= float64(size.KB):
		return size.ParseTargetSize(meta.String(), size.KB).String()
	default:
		return size.ParseTargetSize(meta.String(), size.B).String()
	}
}
