package pgstorage

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is a wrapper providing logging facilities.
type Logger struct {
	x *zap.SugaredLogger
}

func NewLogger(cfg Config) (*Logger, error) {
	var level zap.AtomicLevel
	err := level.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return nil, fmt.Errorf("error on setting log level: %s", err)
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger)),
		level,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()                                     //nolint:gosec,errcheck
	withOptions := logger.WithOptions(zap.AddCallerSkip(2)) //nolint:gomnd
	return &Logger{x: withOptions.Sugar()}, nil
}

// WithFields returns a new Logger with additional fields as per keyValuePairs.
// The original Logger instance is not affected.
func (l *Logger) WithFields(keyValuePairs ...interface{}) *Logger {
	return &Logger{
		x: l.x.With(keyValuePairs...),
	}
}

// Debug calls log.Debug
func (l *Logger) Debug(args ...interface{}) {
	l.x.Debug(args...)
}

// Info calls log.Info
func (l *Logger) Info(args ...interface{}) {
	l.x.Info(args...)
}

// Warn calls log.Warn
func (l *Logger) Warn(args ...interface{}) {
	l.x.Warn(args...)
}

// Error calls log.Error
func (l *Logger) Error(args ...interface{}) {
	l.x.Error(args...)
}

// Fatal calls log.Fatal
func (l *Logger) Fatal(args ...interface{}) {
	l.x.Fatal(args...)
}

// Debugf calls log.Debugf
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.x.Debugf(template, args...)
}

// Infof calls log.Infof
func (l *Logger) Infof(template string, args ...interface{}) {
	l.x.Infof(template, args...)
}

// Warnf calls log.Warnf
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.x.Warnf(template, args...)
}

// Fatalf calls log.Fatalf
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.x.Fatalf(template, args...)
}

// Errorf calls log.Errorf and stores the error message into the ErrorFile
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.x.Errorf(template, args...)
}

// Debugw calls log.Debugw
func (l *Logger) Debugw(msg string, kv ...interface{}) {
	l.x.Debugw(msg, kv...)
}

// Infow calls log.Infow
func (l *Logger) Infow(msg string, kv ...interface{}) {
	l.x.Infow(msg, kv...)
}

// Warnw calls log.Warnw
func (l *Logger) Warnw(msg string, kv ...interface{}) {
	l.x.Warnw(msg, kv...)
}

// Errorw calls log.Errorw
func (l *Logger) Errorw(msg string, kv ...interface{}) {
	l.x.Errorw(msg, kv...)
}

// Fatalw calls log.Fatalw
func (l *Logger) Fatalw(msg string, kv ...interface{}) {
	l.x.Fatalw(msg, kv...)
}
