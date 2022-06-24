package rotzap

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//InitZap Init zap by config and file-rotatelogs
func InitZap(cfg zap.Config, writer io.Writer) (*zap.Logger, error) {
	if cfg.Level == (zap.AtomicLevel{}) {
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	enc := zapcore.NewJSONEncoder(cfg.EncoderConfig)
	if cfg.Encoding == "console" {
		enc = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	} else {
		enc = zapcore.NewJSONEncoder(cfg.EncoderConfig)
	}
	zapLog := zap.New(
		zapcore.NewCore(
			enc,
			zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(writer)),
			cfg.Level,
		),
	)

	//Copy from zap.buildOptions
	var opts []zap.Option
	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zapcore.ErrorLevel
	if cfg.Development {
		stackLevel = zapcore.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	zapLog = zapLog.WithOptions()

	return zapLog, nil
}
