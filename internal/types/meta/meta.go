package meta

import "github.com/dstgo/wilson/pkg/route"

// NoAuth means the router which use this meta has need to authenticate
var NoAuth = route.E{
	Key: "NoAuth",
	Val: struct{}{},
}

func Name(routeName string) route.E {
	return route.E{
		Key: "RouteName",
		Val: routeName,
	}
}

func Comment(routeComment string) route.E {
	return route.E{
		Key: "RouteComment",
		Val: routeComment,
	}
}

func Roles(roles ...string) route.E {
	return route.E{
		Key: "role",
		Val: roles,
	}
}
