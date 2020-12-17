package rotzap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func loadYamlConfig(path string) (*Config, error) {
	cfg := Config{
		Zap: zap.NewDevelopmentConfig(),
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func loadJsonConfig(path string) (*Config, error) {
	cfg := Config{
		Zap: zap.NewDevelopmentConfig(),
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loadRotZapConfig Load config file.
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

	if ext == ".yaml" || ext == ".yml" {
		return loadYamlConfig(path)
	} else if ext == ".json" {
		return loadJsonConfig(path)
	} else {
		if fileExists(path + ".yaml") {
			return loadYamlConfig(path + ".yaml")
		} else if fileExists(path + ".yml") {
			return loadYamlConfig(path + ".yml")
		} else if fileExists(path + ".json") {
			return loadJsonConfig(path + ".json")
		}
	}

	return nil, fmt.Errorf("Config file is not exists!")
}

// InitRotZap init file-rotatelogs and zap by config.
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

// InitRotZapFromYaml init file-rotatelogs and zap from a yaml configuration string.
func InitRotZapFromYaml(yamlStr string) (*zap.Logger, error) {
	cfg := Config{
		Zap: zap.NewDevelopmentConfig(),
	}

	err := yaml.Unmarshal([]byte(yamlStr), &cfg)
	if err != nil {
		return nil, err
	}

	return InitRotZap(cfg.Rot, cfg.Zap)
}

// InitRotZapFromJSON init file-rotatelogs and zap from a json configuration string.
func InitRotZapFromJSON(jsonStr string) (*zap.Logger, error) {
	cfg := Config{
		Zap: zap.NewDevelopmentConfig(),
	}

	err := json.Unmarshal([]byte(jsonStr), &cfg)
	if err != nil {
		return nil, err
	}

	return InitRotZap(cfg.Rot, cfg.Zap)
}

//InitRotZapFromCfgFile Init file-rotatelogs and zap from a config file. Only support 'yaml' and 'json' format.
func InitRotZapFromCfgFile(cfgName string) (*zap.Logger, error) {
	rotZapCfg, err := loadRotZapConfig(cfgName)
	if err != nil {
		return nil, err
	}

	return InitRotZap(rotZapCfg.Rot, rotZapCfg.Zap)
}
