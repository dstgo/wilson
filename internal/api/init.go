package api

import (
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/types/meta"
	roleType "github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/pkg/ginx"
)

func initApiRouterACL(group *ginx.RouterGroup, resolver role.Resolver) error {

	perms := make([]roleType.PermInfo, 0, 10)

	group.Walk(func(info ginx.WalkRouteInfo) error {
		if info.IsGroup {
			return nil
		}

		var (
			name      string
			groupName string
			tag       = system.OpenAPI
		)

		routeName, b := info.Meta.Get(meta.Name("").Key)
		if !b {
			routeName.Val = info.FullPath
		}
		name = routeName.String()

		if info.Group != nil {
			group, b := info.Group.Meta.Get(meta.Group("").Key)
			if !b {
				group.Val = info.Group.FullPath
			}
			groupName = group.String()
		}

		permInfo := roleType.PermInfo{
			Name:   name,
			Object: info.FullPath,
			Group:  groupName,
			Action: info.Method,
			Tag:    tag,
		}

		perms = append(perms, permInfo)

		return nil
	})

	return resolver.CreatePermInBatch(perms)
}
