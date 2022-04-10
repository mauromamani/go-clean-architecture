package logger

import (
	"os"

	"github.com/mauromamani/go-clean-architecture/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type apiLogger struct {
	cfg    *config.Config
	logger *zap.Logger
}

func New(cfg *config.Config) *apiLogger {
	return &apiLogger{
		cfg: cfg,
	}
}

// InitLogger:
func (l *apiLogger) InitLogger() {
	logWritter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.TimeKey = "time"
	encoderCfg.NameKey = "name"
	encoderCfg.MessageKey = "message"

	// Define JSON or Console encoder
	if l.cfg.Logger.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(encoder, logWritter, zapcore.DebugLevel)

	l.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Logger methods
func (l *apiLogger) Debug(args ...interface{}) {
	l.logger.Sugar().Debug(args...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.logger.Sugar().Info(args...)
}

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.logger.Sugar().Warn(args...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *apiLogger) Error(args ...interface{}) {
	l.logger.Sugar().Error(args...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.logger.Sugar().DPanic(args...)
}

func (l *apiLogger) DPanicf(template string, args ...interface{}) {
	l.logger.Sugar().DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.logger.Sugar().Panic(args...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.logger.Sugar().Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.logger.Sugar().Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Sugar().Fatalf(template, args...)
}
