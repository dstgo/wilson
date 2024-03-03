package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/246859/steamapi"
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/assets"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/handler"
	"github.com/dstgo/wilson/internal/handler/middleware"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/ginx/bind"
	"github.com/dstgo/wilson/pkg/sysinfo"
	"github.com/gin-gonic/gin"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	googlerpc "google.golang.org/grpc"
	"io"
	"net/http"
	"path"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// on server boot hooks

func LogBanner(buildInfo conf.BuildInfo, logConf *conf.LogConf, logger *logrus.Logger, bannerPath string) error {
	bannerTemplate := bytes.NewBuffer(nil)

	banner, err := template.ParseFS(assets.Fs, bannerPath)
	if err != nil {
		return err
	}

	hostInfo := sysinfo.GetHostInfo()
	cpuInfo := sysinfo.GetCpuInfo()

	bannerData := map[string]any{
		"author":    buildInfo.Author,
		"version":   buildInfo.Version,
		"buildTime": buildInfo.BuildTime,
		"logMode":   strings.ToUpper(logConf.Level),
		"goVersion": runtime.Version(),
		"osInfo":    fmt.Sprintf("%s %s", hostInfo.Os, hostInfo.Version),
		"timezone":  time.Now().Format("MST -07"),
		"archInfo":  runtime.GOARCH,
		"cpuInfo":   cpuInfo.Name,
	}

	if err := banner.Execute(bannerTemplate, bannerData); err != nil {
		return err
	}

	logger.Infoln(fmt.Sprintf("\n\n%s", bannerTemplate.String()))

	return nil
}

func LoadDataSource(ctx context.Context, dataConf *conf.DataConf) (*data.DataSource, error) {

	log.L().Infoln("attempt to load datasource...")
	datasource, err := data.NewDataSource(ctx, dataConf)
	if err != nil {
		log.L().Errorf("load data datasource failed: %s", err.Error())
		return datasource, err
	}
	log.L().Infof("load data datasource ok √")
	return datasource, nil
}

// on server shutdown hooks

func CloseDataSource(datasource *data.DataSource) {
	if datasource != nil {
		// close datasource
		if err := datasource.Close(); err != nil {
			log.L().Errorf("data source closed failed: %s", err)
			return
		}
	}
	log.L().Infoln("data source closed successfully")
}

func NewLocale(cfg *locale.Conf) (*locale.Locale, error) {
	l, err := locale.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	locale.Setup(l)
	return l, nil
}

// NewLogger config logrus middleware
func NewLogger(logConf *conf.LogConf) (*log.Logger, error) {
	logConf.TimeFormat = types.DateTimeFormat
	logConf.Order = []string{
		types.LogIpKey, types.LogHttpMethodKey, types.LogHttpStatusKey, types.LogRequestCostKey,
		types.LogRequestPathKey, types.LogRequestUrlKey, types.LogRequestQuery, types.LogRequestHeader,
		types.LogRequestContentType, types.LogHttpContentLength, types.LogRequestBody,
		types.LogResponseContentType, types.LogHttpResponseLength,
		types.LogRecoverRequestKey, types.LogRecoverErrorKey,
		types.LogRecoverStackKey, types.LogRequestIdKey}
	logger, err := log.NewLogger(logConf)
	if err != nil {
		return nil, errors.Wrap(err, "load logger failed")
	}
	return logger, nil
}

// NewHttpServer initializes http server configuration
func NewHttpServer(cfg *conf.WilsonConf, lang *locale.Locale, logger *logrus.Logger) (*gin.Engine, *http.Server) {

	serverConf := cfg.ServerConf

	gin.DisableConsoleColor()
	gin.DisableBindValidation()
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	engine := gin.New()

	engine.MaxMultipartMemory = serverConf.HttpConf.MultipartMax

	engine.Use(
		middleware.UseRequestId(),
		middleware.UseLogger(logger, path.Join(handler.DocPath, "*any"), path.Join(api.DocPath, "*any")),
		middleware.UseRecovery(logger),
		middleware.UseAcceptLanguage(lang.Default()),
	)

	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NotFoundHandler())

	bind.HandlerChain = append(bind.HandlerChain, middleware.BindBadParamsHandler())

	server := &http.Server{
		Addr:              serverConf.HttpConf.Address,
		ReadTimeout:       serverConf.HttpConf.ReadTimeout,
		ReadHeaderTimeout: serverConf.HttpConf.ReadHeadTimeout,
		WriteTimeout:      serverConf.HttpConf.WriteTimeout,
		IdleTimeout:       serverConf.HttpConf.IdleTimeout,
		MaxHeaderBytes:    serverConf.HttpConf.MaxHeader,
	}

	server.Handler = engine

	return engine, server
}

// NewGRPCServer initializes a new grpc server with kratos framework
func NewGRPCServer(grpcConf *conf.GrpcConf, logger kratoslog.Logger) *grpc.Server {

	// transport config
	confOptions := grpc.Options(
		googlerpc.MaxRecvMsgSize(grpcConf.MaxRecv),
		googlerpc.MaxSendMsgSize(grpcConf.MaxSend),
		googlerpc.WriteBufferSize(grpcConf.WriteBuffer),
		googlerpc.ReadBufferSize(grpcConf.ReadBuffer),
		googlerpc.MaxHeaderListSize(grpcConf.MaxHeaderSize),
	)

	// grpc middleware
	middlewareOptions := grpc.Middleware(
		// panic recovery
		recovery.Recovery(),
		// request logger
		logging.Server(logger),
		// bbr rate limiting
		ratelimit.Server(),
		// params validator
		validate.Validator(),
	)

	grpcServer := grpc.NewServer(
		grpc.Network("tcp"),
		grpc.Address(grpcConf.Address),
		confOptions,
		middlewareOptions,
	)

	return grpcServer
}

// NewDockerClient returns a DockerClient from local environment
func NewDockerClient(ctx context.Context) (*client.Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	// ping docker daemon to test availability
	if _, err := dockerClient.Ping(ctx); err != nil {
		return nil, err
	}
	log.L().Infoln("local docker client established successfully √")
	return dockerClient, nil
}

func NewSteamClient(cfg *conf.DstConf) (*steamapi.Client, error) {
	steamClient, err := steamapi.New(cfg.SteamKey)
	if err != nil {
		return nil, err
	}
	// test steam web api availability
	if _, err := steamClient.ISteamWebAPIUtil().GetServerInfo(); err != nil {
		return nil, err
	}
	log.L().Infoln("steam webapi client established successfully √")
	return steamClient, nil
}
