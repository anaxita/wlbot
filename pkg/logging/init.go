package logging

import (
	"wlbot/internal/xerrors"
	"wlbot/pkg/version"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(debug bool, logfile string) (*zap.SugaredLogger, error) {
	var (
		zapLevel   = zap.InfoLevel
		encodingAs = "json"
	)

	if debug {
		zapLevel = zap.DebugLevel
		encodingAs = "console"
	}

	const percents = 100

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       debug,
		DisableCaller:     !debug,
		DisableStacktrace: !debug,
		Sampling: &zap.SamplingConfig{
			Hook:       nil,
			Initial:    percents,
			Thereafter: percents,
		},
		Encoding: encodingAs,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "msg",
			LevelKey:            "level",
			TimeKey:             "ts",
			NameKey:             "logger",
			CallerKey:           "caller",
			FunctionKey:         zapcore.OmitKey,
			StacktraceKey:       "stacktrace",
			SkipLineEnding:      false,
			LineEnding:          zapcore.DefaultLineEnding,
			EncodeLevel:         zapcore.LowercaseLevelEncoder,
			EncodeTime:          zapcore.RFC3339TimeEncoder,
			EncodeDuration:      zapcore.SecondsDurationEncoder,
			EncodeCaller:        zapcore.ShortCallerEncoder,
			EncodeName:          nil,
			NewReflectedEncoder: nil,
			ConsoleSeparator:    "",
		},
		OutputPaths:      []string{"stderr", logfile},
		ErrorOutputPaths: []string{"stderr", logfile},
		InitialFields:    map[string]interface{}{"version": version.V},
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, xerrors.Wrap(err, "build logger")
	}

	return l.Sugar(), nil
}
