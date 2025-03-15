package kratosx

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/authentication"
	"github.com/dstgo/wilson/framework/kratosx/library/captcha"
	"github.com/dstgo/wilson/framework/kratosx/library/configure"
	"github.com/dstgo/wilson/framework/kratosx/library/db"
	"github.com/dstgo/wilson/framework/kratosx/library/email"
	"github.com/dstgo/wilson/framework/kratosx/library/http"
	"github.com/dstgo/wilson/framework/kratosx/library/ip"
	"github.com/dstgo/wilson/framework/kratosx/library/jwt"
	"github.com/dstgo/wilson/framework/kratosx/library/loader"
	"github.com/dstgo/wilson/framework/kratosx/library/logger"
	"github.com/dstgo/wilson/framework/kratosx/library/pool"
	"github.com/dstgo/wilson/framework/kratosx/library/prometheus"
	rd "github.com/dstgo/wilson/framework/kratosx/library/redis"
)

type Context interface {
	Env() string
	Logger() *log.Helper
	DB(name ...string) *gorm.DB
	Transaction(fn func(ctx Context) error, name ...string) error
	Redis(name ...string) *redis.Client
	Go(runner pool.Runner) error
	Loader(name string) []byte
	Email() email.Email
	ClientIP() string
	Captcha() captcha.Captcha
	JWT() jwt.Jwt
	Token() string
	Authentication() authentication.Authentication
	Ctx() context.Context
	GetMetadata(string) string
	SetMetadata(key, value string)
	WaitRunner() pool.WaitRunner
	HTTP() *resty.Client
	HTTPRequest() http.Request
	GrpcConn(srvName string) (*grpc.ClientConn, error)

	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Config() config.Config
	Endpoint() []string

	Clone() Context
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

type ctx struct {
	context.Context
	kratos.AppInfo
}

type cloneCtx struct {
	child  context.Context
	parent context.Context
}

// Ctx 获取行下文ctx
func (c *ctx) Ctx() context.Context {
	return c.Context
}

// Logger 获取链路日志器
func (c *ctx) Logger() *log.Helper {
	if !c.Config().IsInit() {
		return log.NewHelper(log.DefaultLogger)
	}
	return logger.Helper().WithContext(c)
}

func (c *ctx) Transaction(fn func(ctx Context) error, name ...string) error {
	dbi := db.Instance()
	if _, ok := c.Value(dbi.TxKey(name...)).(*gorm.DB); ok {
		return fn(MustContext(c.Ctx()))
	}

	return dbi.Get(name...).WithContext(c.Ctx()).Transaction(func(tx *gorm.DB) error {
		cc := context.WithValue(c.Ctx(), dbi.TxKey(name...), tx)
		return fn(MustContext(cc))
	})
}

// DB 数据库实例
func (c *ctx) DB(name ...string) *gorm.DB {
	dbi := db.Instance()
	tx, ok := c.Value(dbi.TxKey(name...)).(*gorm.DB)
	if ok {
		return tx
	}
	return dbi.Get(name...).WithContext(c.Ctx())
}

// Prometheus 监控
func (c *ctx) Prometheus() prometheus.Prometheus {
	return prometheus.Instance()
}

// Redis 获取缓存实例
func (c *ctx) Redis(name ...string) *redis.Client {
	return rd.Instance().Get(name...).WithContext(c.Context)
}

// Go 获取并发池实例
func (c *ctx) Go(runner pool.Runner) error {
	return pool.Instance().Go(runner)
}

// WaitRunner 获取并发池等待实例
func (c *ctx) WaitRunner() pool.WaitRunner {
	return pool.NewWaitRunner()
}

// Loader 获加载器实例
func (c *ctx) Loader(name string) []byte {
	return loader.Instance().Get(name)
}

// Email 获取邮箱实例
func (c *ctx) Email() email.Email {
	return email.Instance()
}

// Captcha 获取图形验证器
func (c *ctx) Captcha() captcha.Captcha {
	return captcha.Instance()
}

// ClientIP 获取客户端IP地址
func (c *ctx) ClientIP() string {
	return ip.ClientIP(c.Context)
}

// JWT 获取令牌验证器
func (c *ctx) JWT() jwt.Jwt {
	return jwt.Instance()
}

func (c *ctx) HTTP() *resty.Client {
	return http.Instance()
}

// HTTPRequest 带链路日志的请求工具
func (c *ctx) HTTPRequest() http.Request {
	cfg := c.Config().App().Http
	cfg.Server = c.Name()
	return http.NewRequest(cfg, c.Logger())
}

// Token 获取令牌验证器
func (c *ctx) Token() string {
	return jwt.Instance().GetToken(c.Context)
}

// Authentication 获取权限验证器
func (c *ctx) Authentication() authentication.Authentication {
	return authentication.Instance()
}

// GetMetadata 获取元数据信息
func (c *ctx) GetMetadata(key string) string {
	if values, ok := metadata.FromServerContext(c.Context); ok {
		return values.Get(key)
	}
	return ""
}

// SetMetadata 设置元数据信息
func (c *ctx) SetMetadata(key, value string) {
	c.Context = md.AppendToClientContext(c.Context, key, value)
}

// Config 获取配置对象
func (c *ctx) Config() config.Config {
	return config.Instance()
}

// GrpcConn 获取grpc 连接具柄
func (c *ctx) GrpcConn(srvName string) (*grpc.ClientConn, error) {
	cli := configure.Get(srvName)
	if cli == nil {
		return nil, errors.New("not exist server " + srvName)
	}
	return cli.Conn(c.Ctx())
}

func (c *ctx) Clone() Context {
	return MustContext(context.WithoutCancel(c.Context))
}

// TraceID 获取trace id
func (c *ctx) TraceID() string {
	t, _ := tracing.TraceID()(c.Context).(string)
	return t
}

// SpanID 获取span id
func (c *ctx) SpanID() string {
	t, _ := tracing.SpanID()(c.Context).(string)
	return t
}

// Env 获取配置环境
func (c *ctx) Env() string {
	return c.Config().App().Env
}

func (c *ctx) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *ctx) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *ctx) Err() error {
	return c.Context.Err()
}

func (c *ctx) Value(key any) any {
	return c.Context.Value(key)
}

func (c *cloneCtx) Ctx() context.Context {
	return c
}

func (c *cloneCtx) Deadline() (deadline time.Time, ok bool) {
	return c.child.Deadline()
}

func (c *cloneCtx) Done() <-chan struct{} {
	return c.child.Done()
}

func (c *cloneCtx) Err() error {
	return c.child.Err()
}

func (c *cloneCtx) Value(key any) any {
	return c.parent.Value(key)
}
