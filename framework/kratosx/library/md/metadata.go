package md

import (
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/valx"

	"github.com/dstgo/wilson/api/gen/errors"
)

type Auth struct {
	UserId            uint32 `json:"userId"`
	RoleId            uint32 `json:"roleId"`
	RoleKeyword       string `json:"roleKeyword"`
	DepartmentId      uint32 `json:"departmentId"`
	DepartmentKeyword string `json:"departmentKeyword"`
}

func NewAuthMap(info *Auth) map[string]any {
	var res map[string]any
	_ = valx.Transform(info, &res)
	return res
}

func GetAuthInfo(ctx kratosx.Context) *Auth {
	var (
		data Auth
		err  error
	)
	if ctx.Token() != "" {
		err = ctx.JWT().Parse(ctx, &data)
	} else {
		// 三方服务调用的时候通过auth信息获取
		err = ctx.Authentication().ParseAuthFromMD(ctx, &data)
	}
	if err != nil {
		panic(errors.ForbiddenErrorWrap(err))
	}

	if data.UserId == 0 {
		panic(errors.ForbiddenErrorWrap(err))
	}
	return &data
}

func UserId(ctx kratosx.Context) uint32 {
	return GetAuthInfo(ctx).UserId
}

func RoleId(ctx kratosx.Context) uint32 {
	return GetAuthInfo(ctx).RoleId
}

func RoleKeyword(ctx kratosx.Context) string {
	return GetAuthInfo(ctx).RoleKeyword
}

func DepartmentId(ctx kratosx.Context) uint32 {
	return GetAuthInfo(ctx).DepartmentId
}

func DepartmentKeyword(ctx kratosx.Context) string {
	return GetAuthInfo(ctx).DepartmentKeyword
}

const (
	ServiceAppName    = "service:app_name"
	ServiceAppVersion = "service:app_version"
	ServiceBuildTime  = "service:build_time"
)

func AppName(ctx kratosx.Context) string {
	return ctx.GetMetadata(ServiceAppName)
}

func AppVersion(ctx kratosx.Context) string {
	return ctx.GetMetadata(ServiceAppVersion)
}

func BuildTime(ctx kratosx.Context) string {
	return ctx.GetMetadata(ServiceBuildTime)
}
