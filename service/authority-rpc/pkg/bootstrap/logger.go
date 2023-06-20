package bootstrap

import (
	"authority-rpc/internal/conf"
	zapLogger "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

const DefaultLogPath = "./logs"

func NewLoggerProvider(cfg *conf.Logger, serviceInfo *ServiceInfo) log.Logger {
	l := NewZapLogger(cfg)

	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		//"time", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// NewZapLogger 创建一个新的日志记录器 - Zap
func NewZapLogger(cfg *conf.Logger) log.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(cfg.Zap.FilePath); !exist {
		if cfg.Zap.FileName == "" {
			cfg.Zap.FilePath = DefaultLogPath
		}
		if err := os.MkdirAll(cfg.Zap.FilePath, os.ModePerm); err != nil {
			cfg.Zap.FilePath = DefaultLogPath
			if err := os.MkdirAll(cfg.Zap.FilePath, os.ModePerm); err != nil {
				panic(err)
			}
		}
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(cfg.Zap.FilePath, cfg.Zap.FileName), // 日志文件路径
		MaxSize:    int(cfg.Zap.MaxSize),
		MaxBackups: int(cfg.Zap.MaxBackups),
		MaxAge:     int(cfg.Zap.MaxAge),
	}

	var writeSyncer zapcore.WriteSyncer

	writeSyncer = zapcore.AddSync(lumberJackLogger)

	if cfg.Zap.LogStdout {
		// 日志同时输出到控制台和日志文件中
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	}

	var lvl = new(zapcore.Level)
	if err := lvl.UnmarshalText([]byte(cfg.Zap.Level)); err != nil {
		return nil
	}

	core := zapcore.NewCore(jsonEncoder, writeSyncer, lvl)
	logger := zap.New(core).WithOptions()

	wrapped := zapLogger.NewLogger(logger)

	return wrapped
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
