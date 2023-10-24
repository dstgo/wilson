package test

import (
	"fmt"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHttpRoute(t *testing.T) {
	engine := gin.Default()
	group := ginx.NewRouterGroup(&engine.RouterGroup)
	group.GET("/ping", nil)
	userGroup := group.Group("/user", nil)
	userGroup.GET("/info", nil)

	group.Walk(func(info ginx.WalkRouteInfo) error {
		fmt.Println(info)
		return nil
	})
}

func TestHttpQualityValues(t *testing.T) {
	// test1 empty
	{
		accept_language := ""
		langs := httpx.GetQualityValuePairs(accept_language)
		// []
		t.Log(langs)
	}
	// test2 single lang
	{
		accept_language := "zh-CN"
		langs := httpx.GetQualityValuePairs(accept_language)
		// [zh-CN]
		t.Log(langs)
	}
	// test3 multi lang
	{
		accept_language := "zh-CN,en-US,jp"
		langs := httpx.GetQualityValuePairs(accept_language)
		// [zh-CN en-US jp]
		t.Log(langs)
	}
	// test4 quality
	{
		accept_language := "en-US,en;q=0.5"
		langs := httpx.GetQualityValuePairs(accept_language)
		// [zh-CN en-US jp]
		t.Log(langs)
	}
	// test5 multi lang and quality
	{
		accept_language := "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"
		langs := httpx.GetQualityValuePairs(accept_language)
		// [zh-CN zh zh-TW zh-HK en-US en]
		t.Log(langs)
	}
	// test6
	{
		accept_language := "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
		langs := httpx.GetQualityValuePairs(accept_language)
		// [text/html application/xhtml+xml application/xml */*]
		t.Log(langs)
	}
}

func TestBearerToken(t *testing.T) {
	// test1
	{
		header := "Bearer 123456"
		t.Log(httpx.GetBearerToken(header))
	}
	// test2
	{
		header := "Bearer asdasjljzda"
		t.Log(httpx.GetBearerToken(header))
	}
	// test3
	{
		header := "Bearer "
		t.Log(httpx.GetBearerToken(header))
	}
	// test4
	{
		header := "Bearerasd "
		t.Log(httpx.GetBearerToken(header))
	}
	// test5
	{
		header := ""
		t.Log(httpx.GetBearerToken(header))
	}
}
