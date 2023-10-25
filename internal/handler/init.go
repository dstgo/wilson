package handler

import (
	"github.com/dstgo/wilson/internal/core/log"
	roleSo "github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/internal/types/role"
	usert "github.com/dstgo/wilson/internal/types/user"
	"github.com/dstgo/wilson/pkg/ginx"
	"reflect"
)

func initFirstUser(source *data.DataSource) (entity.User, error) {
	info := user.NewUserInfo(source)

	count, err := user.Count(source.ORM())
	if err != nil {
		return entity.User{}, err
	}

	if count == 0 {
		modify := user.NewUserModify(source, info)
		err := modify.Create(usert.InitialUser)
		if err != nil {
			return entity.User{}, err
		}
		userEn, err := user.GetUserByName(source.ORM(), usert.InitialUser.Username)
		if err != nil {
			return entity.User{}, err
		}
		return userEn, nil
	}

	return entity.User{}, nil
}

func initRouterRole(source *data.DataSource, root *ginx.RouterGroup, resolver roleSo.Resolver) error {

	log.L().Info("starting to initialize router acl...")
	var (
		permsMap = make(map[string][]role.PermInfo)
	)

	err := resolver.CreateRoleInBatch([]role.RoleInfo{
		role.AdminRole,
		role.UserRole,
		role.AnonymousRole,
	})

	if err != nil {
		return err
	}

	root.Walk(func(info ginx.WalkRouteInfo) error {
		if info.IsGroup {
			return nil
		}

		var (
			name      string
			groupName string
			tag       = "appapi"
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

		permInfo := role.PermInfo{
			Name:   name,
			Object: info.FullPath,
			Group:  groupName,
			Action: info.Method,
			Tag:    tag,
		}

		// must be []string
		roles, b := info.Meta.Get(meta.Roles().Key)
		if roles.Val == nil {
			return nil
		}

		for rs, i := reflect.ValueOf(roles.Val), 0; i < rs.Len(); i++ {
			r := rs.Index(i).Interface().(role.RoleInfo)
			permsMap[r.Code] = append(permsMap[r.Code], permInfo)

		}

		return nil
	})

	// related role and permissions
	for roleCode, perms := range permsMap {
		if err := resolver.CreateRolePermBatch(role.RoleInfo{Code: roleCode}, perms); err != nil {
			return err
		}
	}

	// initialize first user roles
	firstUser, err := initFirstUser(source)
	if err != nil {
		return err
	}

	// if first user created, then
	if firstUser.Id > 0 {
		userRole := user.NewUserModify(source, user.NewUserInfo(source))
		err := userRole.SaveRolesByCode(firstUser.UUID, []string{
			role.AdminRole.Code,
			role.UserRole.Code,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
