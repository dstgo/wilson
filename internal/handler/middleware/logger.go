package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dstgo/size"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"strings"
	"time"
)

func UseRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := uuid.NewString()
		// set response header X-Request-ID
		httpx.SetRequestId(ctx, requestId)
		ctx.Next()
	}
}

func UseLogger(logger *logrus.Logger, skips ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// skip logger
		if slices.Contains(skips, ctx.FullPath()) {
			ctx.Next()
			return
		}

		var (
			request   = ctx.Request
			requestId = httpx.GetRequestId(ctx)
			startTime = time.Now()
		)

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
			contentType  = httpx.GetRequestContentType(ctx)
			responseType = httpx.GetResponseContentType(ctx)

			body  string
			query string

			clientIp = ctx.ClientIP()
			err      = ctx.Errors
		)

		query = describeJson(ctx, ctx.Request.URL.Query())

		switch ctx.ContentType() {
		case gin.MIMEJSON:
			rawJson, err := httpx.DescribeRawJson(ctx.Request.Body)
			if err != nil {
				ctx.Error(err)
			}
			body = string(rawJson)
		case gin.MIMEMultipartPOSTForm, gin.MIMEPOSTForm:
			form, err := ctx.MultipartForm()
			if err != nil {
				ctx.Error(err)
			} else {
				body = describeJson(ctx, httpx.DescribeFormData(form))
			}
		}

		if len(path) == 0 {
			path = "not found"
		}

		entry := logger.WithContext(ctx).
			// IP
			WithField(types.LogIpKey, clientIp).
			// method status cost
			WithField(types.LogHttpMethodKey, method).WithField(types.LogHttpStatusKey, status).WithField(types.LogRequestCostKey, fmt.Sprintf("%dms", costTime)).
			// path url query
			WithField(types.LogRequestPathKey, path).WithField(types.LogRequestUrlKey, url).WithField(types.LogRequestQuery, query).
			// request type length body
			WithField(types.LogRequestContentType, contentType).WithField(types.LogHttpContentLength, closedSize(requestSize)).WithField(types.LogRequestBody, body).
			// response type length
			WithField(types.LogResponseContentType, responseType).WithField(types.LogHttpResponseLength, closedSize(int64(responseSize))).
			// trace id
			WithField(types.LogRequestIdKey, requestId)

		// log by status
		switch {
		case err == nil && 100 < status && status < 400:
			entry.Infoln()
		case status >= 400:
			entry.
				WithError(errors.New(strings.Join(err.Errors(), ";"))).
				Errorln()
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

func describeJson(ctx *gin.Context, val any) string {
	marshal, err := json.Marshal(val)
	if err != nil {
		ctx.Error(err)
	}
	return string(marshal)
}
