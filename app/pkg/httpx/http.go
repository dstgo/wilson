package httpx

import (
	"github.com/dstgo/wilson/app/pkg/httpx/httpheader"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"regexp"
	"sort"
	"strings"
)

var (
	qualityValueRegexp = regexp.MustCompile("[/+*\\w-;q=.]+")
)

func SetRequestId(ctx *gin.Context, id string) {
	ctx.Set(httpheader.RequestIdHeader, id)
	ctx.Writer.Header().Set(httpheader.RequestIdHeader, id)
}

func GetRequestId(ctx *gin.Context) (requestId string) {
	return ctx.GetString(httpheader.RequestIdHeader)
}

// GetAcceptLanguage
// the header should compliance with regulations follow
// Accept-Language: de
// Accept-Language: de-CH
// Accept-Language: en-US,en;q=0.5
// Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept-Language
func GetAcceptLanguage(ctx *gin.Context) []string {
	cached := ctx.GetStringSlice(httpheader.AcceptLanguage)
	if cached != nil {
		return cached
	}
	qualityValuePairs := GetQualityValuePairs(ctx.GetHeader(httpheader.AcceptLanguage))
	ctx.Set(httpheader.AcceptLanguage, qualityValuePairs)
	return qualityValuePairs
}
func GetFirstAcceptLanguage(ctx *gin.Context) string {
	language := GetAcceptLanguage(ctx)
	if len(language) > 0 {
		return language[0]
	}
	return ""
}

type qualityV struct {
	value   string
	quality float64
}

// GetQualityValuePairs return http quality value pairs ordered by quality from string
// eg: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
//
//	-> [text/html application/xhtml+xml application/xml */*]
//
// reference https://developer.mozilla.org/zh-CN/docs/Glossary/Quality_values
func GetQualityValuePairs(header string) []string {

	var (
		httpLang []string
		qvs      []qualityV
		tempqvs  []qualityV
	)

	allString := qualityValueRegexp.FindAllString(header, -1)

	for _, str := range allString {
		qv := qualityV{
			value:   "",
			quality: 1,
		}

		qvSplit := strings.Split(str, ";q=")
		if len(qvSplit) > 0 {
			qv.value = strings.TrimSpace(qvSplit[0])
			tempqvs = append(tempqvs, qv)
		}

		if len(qvSplit) > 1 {
			quality := cast.ToFloat64(strings.TrimSpace(qvSplit[1]))
			for i := 0; i < len(tempqvs); i++ {
				tempqvs[i].quality = quality
			}
			qvs = append(qvs, tempqvs...)
			tempqvs = []qualityV{}
		}
	}

	if len(tempqvs) > 0 {
		qvs = append(qvs, tempqvs...)
	}

	// sort by quality
	sort.Slice(qvs, func(i, j int) bool {
		return qvs[i].quality > qvs[j].quality
	})

	for _, qv := range qvs {
		httpLang = append(httpLang, qv.value)
	}

	return httpLang
}

// GetBearerTokenFromCtx get bearer token from Authorization Header
// param ctx *gin.Context
// return string
func GetBearerTokenFromCtx(ctx *gin.Context) string {
	return GetBearerToken(ctx.GetHeader(httpheader.Authorization))
}

func GetBearerToken(authHeader string) string {
	var token string
	if len(authHeader) == 0 {
		return token
	}
	if !strings.HasPrefix(authHeader, httpheader.BearerToken) {
		return token
	}
	return strings.TrimSpace(strings.TrimPrefix(authHeader, httpheader.BearerToken))
}

func GetContentType(ctx *gin.Context) string {
	return ctx.GetHeader(httpheader.ContentType)
}
