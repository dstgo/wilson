package meta

import (
	"github.com/dstgo/wilson/pkg/ginx"
)

// NoAuth means the router which use this meta has no need to authenticate
var NoAuth = ginx.E{
	Key: "NoAuth",
	Val: struct{}{},
}

// Anonymous means the router which use this meta has no need to authorize
var Anonymous = ginx.E{
	Key: "Anonymous",
	Val: struct{}{},
}

func Name(routeName string) ginx.E {
	return ginx.E{
		Key: "RouteName",
		Val: routeName,
	}
}

func Group(routeName string) ginx.E {
	return ginx.E{
		Key: "GroupRouteName",
		Val: routeName,
	}
}

func Comment(routeComment string) ginx.E {
	return ginx.E{
		Key: "RouteComment",
		Val: routeComment,
	}
}

func Roles(roles ...string) ginx.E {
	return ginx.E{
		Key: "roles",
		Val: roles,
	}
}
