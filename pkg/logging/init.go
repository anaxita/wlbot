package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(debug bool, logfile string) (*zap.SugaredLogger, error) {
	var zapLevel = zap.InfoLevel
	var encodingAs = "json"

	if debug {
		zapLevel = zap.DebugLevel
		encodingAs = "console"
	}

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       debug,
		DisableCaller:     !debug,
		DisableStacktrace: !debug,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: encodingAs,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr", logfile},
		ErrorOutputPaths: []string{"stderr", logfile},
		InitialFields:    nil,
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return l.Sugar(), nil
}
