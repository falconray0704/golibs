package logging

import (
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//go:generate mockgen -destination=mock_logging/logger.go . Logger
type Logger interface {
	// Sync synchronizes the buffer to ensure everything is printed out
	Sync() error
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	// WithPrefix returns another logger that will use the prefix given for all logging operations
	WithPrefix(prefix string) Logger
}

type logger struct {
	prefix    string
	color     *color.Color
	zapLogger Zap
}

func NewLogger(encoding Encoding, level Level, nodeID int) (l Logger, err error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(string(level))); err != nil {
		return nil, err
	}
	zapLogger, err := zap.Config{
		Level:    zap.NewAtomicLevelAt(zapLevel),
		Encoding: string(encoding),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}.Build()
	if err != nil {
		return nil, fmt.Errorf("logger initialization failed: %w", err)
	}
	if nodeID != -1 {
		zapLogger = zapLogger.With(zap.Int("node_id", nodeID))
	}
	return &logger{zapLogger: zapLogger}, nil
}

func (l *logger) Sync() error {
	return l.zapLogger.Sync()
}

// NewEmptyLogger returns a logging that does not print anything
func NewEmptyLogger() (l Logger, err error) {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.OutputPaths, loggerConfig.ErrorOutputPaths = nil, nil
	zapLogger, err := loggerConfig.Build()
	if err != nil {
		return nil, err
	}
	return &logger{zapLogger: zapLogger}, nil
}

func (l *logger) WithPrefix(prefix string) Logger {
	return &logger{
		prefix:    l.prefix + prefix,
		zapLogger: l.zapLogger,
	}
}

func (l *logger) WithColor(color *color.Color) Logger {
	return &logger{
		color: color,
	}
}
