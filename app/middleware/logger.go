package middleware

import (
	"fmt"
	"github.com/dstgo/size"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"time"
)

func UseLogger(logger *logrus.Logger, skips ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// skip logger
		if slices.Contains(skips, ctx.FullPath()) {
			ctx.Next()
			return
		}

		var (
			request   = ctx.Request
			requestId = uuid.NewString()
			startTime = time.Now()
		)

		// set response header X-Request-ID
		httpx.SetRequestId(ctx, requestId)

		// next handlers
		ctx.Next()

		var (
			url          = ctx.Request.URL.String()
			path         = ctx.FullPath()
			method       = request.Method
			costTime     = time.Now().Sub(startTime).Milliseconds()
			status       = ctx.Writer.Status()
			responseSize = ctx.Writer.Size()
			requestSize  = request.ContentLength
			contentType  = request.Header.Get("Content-Type")
			responseType = ctx.Writer.Header().Get("Content-Type")

			clientIp = ctx.ClientIP()
			err      = ctx.Err()
		)

		if len(path) == 0 {
			path = "not found"
		}

		entry := logger.WithContext(ctx).
			WithField(types.LogIpKey, clientIp).
			WithField(types.LogHttpMethodKey, method).
			WithField(types.LogRequestPathKey, path).
			WithField(types.LogRequestUrlKey, url).
			WithField(types.LogHttpStatusKey, status).
			WithField(types.LogRequestCostKey, fmt.Sprintf("%dms", costTime)).
			WithField(types.LogRequestContentType, contentType).
			WithField(types.LogHttpContentLength, closedSize(requestSize)).
			WithField(types.LogResponseContentType, responseType).
			WithField(types.LogHttpResponseLength, closedSize(int64(responseSize))).
			WithField(types.LogRequestIdKey, requestId)

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

func closedSize(s int64) string {
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
