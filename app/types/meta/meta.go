package meta

import "github.com/dstgo/wilson/pkg/route"

// NoAuth means the router which use this meta has no authentication
var NoAuth = route.E{
	Key: "NoAuth",
	Val: struct{}{},
}
