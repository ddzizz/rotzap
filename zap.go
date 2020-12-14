package rotzap

import (
	"io"

	"go.uber.org/zap"
)

func InitZap(cfg zap.Config, writer io.Writer) (*zap.Logger, error) {
	zapLog, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return zapLog, nil
}
