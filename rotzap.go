package rotzap

import "go.uber.org/zap"

type RotZapConfig struct {
}

func InitRotZap(rotCfg RotConfig, zapCfg zap.Config) (*zap.Logger, error) {
	rot, err := InitRot(&rotCfg)
	if err != nil {
		return nil, err
	}

	zapLog, err := InitZap(zapCfg, rot)
	if err != nil {
		return nil, err
	}

	return zapLog, nil
}
