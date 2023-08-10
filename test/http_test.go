package test

import (
	"fmt"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHttpRoute(t *testing.T) {
	router := route.NewRouter(gin.Default())
	var NoAuth = route.E{
		K: "noauth",
		V: struct{}{},
	}
	router.GET("ping", route.Metas(NoAuth), nil)
	group := router.Group("a", nil)
	group.GET("abc", nil, nil)

	router.Walk(func(info route.RouterInfo) error {
		fmt.Println(fmt.Sprintf("%+v", info))
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
