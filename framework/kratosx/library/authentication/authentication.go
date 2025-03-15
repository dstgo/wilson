package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"

	rediswatcher "github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-redis/redis/v8"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/db"
	rd "github.com/dstgo/wilson/framework/kratosx/library/redis"
	"github.com/dstgo/wilson/framework/pkg/strs"
)

type Authentication interface {
	AddWhitelist(path string, method string)
	RemoveWhitelist(path, method string)
	IsWhitelist(path string, method string) bool
	Auth(role, path, method string) bool
	GetRole(ctx context.Context) (string, error)
	Enforce() *casbin.Enforcer
	IsSkipRole(role string) bool
	SetAuth(req *http.Request, data any) error
	SetAuthMD(ctx context.Context, data any) (context.Context, error)
	ParseAuthFromMD(ctx context.Context, dst any) error
}

type authentication struct {
	conf     *config.Authentication
	enforcer *casbin.Enforcer
	redis    *redis.Client
	roleKey  string
	skipRole map[string]struct{}
	mutex    sync.RWMutex
}

var instance *authentication

const (
	redisKey  = "rbac_authentication"
	authMdKey = "x-md-global-auth"
)

func Instance() Authentication {
	return instance
}

func Init(conf *config.Authentication, watcher config.Watcher) {
	if conf == nil {
		return
	}

	dbi := db.Instance().Get(conf.DB)
	rdi := rd.Instance().Get(conf.Redis)

	if dbi == nil {
		panic("authentication init error not exist database " + conf.DB)
	}

	if rdi == nil {
		panic("authentication init error not exist redis " + conf.DB)
	}

	// 初始化监听器
	w, err := rediswatcher.NewWatcher(
		rdi.Options().Addr,
		rediswatcher.Username(rdi.Options().Username),
		rediswatcher.Password(rdi.Options().Password),
	)
	if err != nil {
		panic("authentication init watcher error:" + err.Error())
	}

	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := adapter.NewAdapterByDB(db.Instance().Get(conf.DB))
	if err != nil {
		panic("authentication init error:" + err.Error())
	}
	object, _ := casbin.NewEnforcer(m, a)
	if err = object.LoadPolicy(); err != nil {
		panic("authentication init error:" + err.Error())
	}

	// 设置监听器
	_ = object.SetWatcher(w)
	_ = w.SetUpdateCallback(func(s string) {
		_ = object.LoadPolicy()
		log.Errorf("casbin watch load policy")
	})

	instance = &authentication{
		enforcer: object,
		redis:    rdi,
		roleKey:  conf.RoleKey,
		skipRole: make(map[string]struct{}),
		conf:     conf,
	}
	instance.initSkipRole(conf.SkipRole)
	instance.initWhitelist(conf.Whitelist)

	whs := map[string]bool{}
	watcher("authentication.whitelist", func(value config.Value) {
		if err := value.Scan(&whs); err != nil {
			log.Errorf("Authentication Whitelist 配置变更失败：%s", err.Error())
			return
		}
		instance.initWhitelist(whs)
	})

	skips := make([]string, 0)
	watcher("authentication.whitelist", func(value config.Value) {
		if err := value.Scan(&skips); err != nil {
			log.Errorf("Authentication SkipRole 配置变更失败：%s", err.Error())
			return
		}
		instance.initSkipRole(skips)
	})
}

func (a *authentication) initWhitelist(whs map[string]bool) {
	for path, is := range whs {
		arr := strings.Split(path, ":")
		if len(arr) != 2 {
			continue
		}
		if is {
			instance.AddWhitelist(arr[1], arr[0])
		} else {
			instance.RemoveWhitelist(arr[1], arr[0])
		}
	}
}

func (a *authentication) initSkipRole(skips []string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for _, role := range skips {
		a.skipRole[role] = struct{}{}
	}
}

func (a *authentication) SetAuth(req *http.Request, data any) error {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		marshal, err := json.Marshal(data)
		if err != nil {
			return errors.New("auth info format error:" + err.Error())
		}
		req.Header.Set(authMdKey, strs.BytesToString(marshal))
	default:
		req.Header.Set(authMdKey, fmt.Sprintf("%s", data))
	}

	return nil
}

func (a *authentication) SetAuthMD(ctx context.Context, data any) (context.Context, error) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Pointer:
		return a.SetAuthMD(ctx, reflect.ValueOf(data).Elem().Interface())
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		marshal, err := json.Marshal(data)
		if err != nil {
			return nil, errors.New("auth info format error:" + err.Error())
		}
		return metadata.AppendToClientContext(ctx, authMdKey, strs.BytesToString(marshal)), nil
	default:
		return metadata.AppendToClientContext(ctx, authMdKey, fmt.Sprintf("%s", data)), nil
	}
}

func (a *authentication) ParseAuthFromMD(ctx context.Context, dst any) error {
	if md, ok := metadata.FromServerContext(ctx); ok {
		body := md.Get(authMdKey)
		if err := json.Unmarshal([]byte(body), dst); err != nil {
			return errors.New("auth info format error:" + err.Error())
		}
		return nil
	}
	return errors.New("not exist auth info")
}

func (a *authentication) IsSkipRole(role string) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	_, is := a.skipRole[role]
	return is
}

func (a *authentication) path(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

func (a *authentication) AddWhitelist(path string, method string) {
	a.redis.HSet(context.Background(), redisKey, a.path(path, method), 1)
}

func (a *authentication) RemoveWhitelist(path, method string) {
	a.redis.HDel(context.Background(), redisKey, a.path(path, method))
}

func (a *authentication) IsWhitelist(path, method string) bool {
	if !a.conf.EnableGrpc && method == "GRPC" {
		return true
	}
	is, _ := a.redis.HGet(context.Background(), redisKey, a.path(path, method)).Bool()
	return is
}

func (a *authentication) Auth(role, path, method string) bool {
	if a.IsWhitelist(path, method) {
		return true
	}

	// 进行鉴权
	is, _ := a.enforcer.Enforce(role, path, method)
	return is
}

func (a *authentication) GetRole(ctx context.Context) (string, error) {
	tokenInfo, is := kratosJwt.FromContext(ctx)
	if !is { // 跳过jwt白名单的也不检测
		return "", nil
	}
	claims, is := tokenInfo.(jwtv5.MapClaims)
	if !is {
		return "", errors.New("token format error")
	}

	role, is := claims[a.roleKey].(string)
	if !is {
		return "", fmt.Errorf("not exist role field %v", a.roleKey)
	}
	return role, nil
}

func (a *authentication) Enforce() *casbin.Enforcer {
	return a.enforcer
}
