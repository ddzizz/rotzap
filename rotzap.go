package rotzap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kardianos/osext"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config is configuration for rotzap
type Config struct {
	Rot RotConfig  `yaml:"rot" json:"rot"`
	Zap zap.Config `yaml:"zap" json:"zap"`
}

//loadRotZapConfig Load config file.
func loadRotZapConfig(cfgName string) (*Config, error) {
	path := cfgName
	ext := filepath.Ext(cfgName)
	if !filepath.IsAbs(path) {
		exePath, err := osext.ExecutableFolder()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(exePath, path)
	}

	cfg := Config{
		Zap: zap.NewDevelopmentConfig(),
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if ext == ".yml" || ext == ".yaml" {
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	} else if ext == ".json" {
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	}

	return nil, fmt.Errorf("Rotzap config file only support 'yaml' and 'json' format")
}

//InitRotZap Init file-rotatelogs and zap by config.
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

//InitRotZapFromCfgFile Init file-rotatelogs and zap from a config file. Only support 'yaml' and 'json' format.
func InitRotZapFromCfgFile(cfgName string) (*zap.Logger, error) {
	rotZapCfg, err := loadRotZapConfig(cfgName)
	if err != nil {
		return nil, err
	}

	return InitRotZap(rotZapCfg.Rot, rotZapCfg.Zap)
}
