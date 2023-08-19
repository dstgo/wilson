package meta

import "github.com/dstgo/wilson/pkg/route"

// NoAuth means the router which use this meta has need to authenticate
var NoAuth = route.E{
	Key: "NoAuth",
	Val: struct{}{},
}
