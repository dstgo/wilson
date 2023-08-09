package test

import (
	"fmt"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHttp(t *testing.T) {
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
