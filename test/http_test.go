package test

import (
	"github.com/dstgo/wilson/pkg/coco/route"
	"testing"
)

func TestHttp(t *testing.T) {
	router := route.NewRouter(nil)
	var NoAuth = route.E{
		K: "noauth",
		V: struct{}{},
	}
	router.GET("/ping", route.Metas(NoAuth), nil)
}
