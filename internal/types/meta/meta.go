package meta

import "github.com/dstgo/wilson/pkg/route"

// NoAuth means the router which use this meta has need to authenticate
var NoAuth = route.E{
	Key: "NoAuth",
	Val: struct{}{},
}

func Roles(roles ...string) route.E {
	return route.E{
		Key: "role",
		Val: roles,
	}
}
