package logger

import (
	"os"

	kratosZap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/pkg/timex"
)

type LogField map[string]any

type logger struct {
	zap *zap.Logger
	fs  []any
}

var ins *logger

func Instance(opts ...Option) log.Logger {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	zapLog := ins.zap.WithOptions(zap.AddCallerSkip(o.callerSkip))
	return log.With(kratosZap.NewLogger(zapLog), ins.fs...)
}

func Helper(opts ...Option) *log.Helper {
	return log.NewHelper(Instance(opts...))
}

// Init 初始化日志器
func Init(lc *config.Logger, watcher config.Watcher, fields LogField) {
	// 没配置则跳过
	if lc == nil {
		return
	}

	// log field 转换
	var fs []any
	for key, val := range fields {
		fs = append(fs, key, val)
	}

	// 初始化
	ins = &logger{}
	ins.initFactory(lc, fs)

	watcher("log", func(value config.Value) {
		if err := value.Scan(lc); err != nil {
			log.Errorf("watch log config failed: %s", err.Error())
			return
		}
		log.Infof("watch log config successfully")
		// 变更初始化
		ins.initFactory(lc, fs)
	})
}

func (l *logger) initFactory(conf *config.Logger, fs []any) {
	// 创建zap logger
	l.zap = l.newZapLogger(conf)
	l.fs = fs

	gLog := log.With(kratosZap.NewLogger(l.zap), fs...)
	// 设置全局logger
	log.SetLogger(gLog)
}

func (l *logger) newZapLogger(conf *config.Logger) *zap.Logger {
	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(timex.DateTimeMillFormat),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 输出器配置
	var output []zapcore.WriteSyncer
	for _, val := range conf.Output {
		switch val {
		case "stdout":
			output = append(output, zapcore.AddSync(os.Stdout))
		case "file":
			output = append(output, zapcore.AddSync(&lumberjack.Logger{
				Filename:   conf.File.Name,
				MaxSize:    conf.File.MaxSize,
				MaxBackups: conf.File.MaxBackup,
				MaxAge:     conf.File.MaxAge,
				Compress:   conf.File.Compress,
				LocalTime:  conf.File.LocalTime,
			}))
		}
	}

	var encoder zapcore.Encoder
	switch conf.Encode {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		panic("invalid log encode: " + conf.Encode)
	}

	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		panic("invalid log level: " + conf.Level)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(output...),
		level,
	)

	// 添加回调
	var zapOptions []zap.Option
	if conf.Caller {
		callerSkip := 3
		if conf.CallerSkip != nil {
			callerSkip = *conf.CallerSkip
		}
		zapOptions = append(zapOptions, zap.AddCaller())
		zapOptions = append(zapOptions, zap.AddCallerSkip(callerSkip))
	}

	return zap.New(core, zapOptions...)
}
