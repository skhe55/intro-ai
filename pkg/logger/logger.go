package logger

import (
	"intro-ai/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type apiLoger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

func NewApiLogger(cfg *config.Config) *apiLoger {
	return &apiLoger{cfg: cfg}
}

var loggerLevelMap = map[string]zapcore.Level{
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func (l *apiLoger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *apiLoger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	encoderCfg := zap.NewDevelopmentEncoderConfig()

	var encoder zapcore.Encoder

	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

func (l *apiLoger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *apiLoger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *apiLoger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *apiLoger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *apiLoger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *apiLoger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}
