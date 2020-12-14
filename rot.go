package rotzap

import (
	"path/filepath"
	"time"

	"github.com/kardianos/osext"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type RotConfig struct {
	Path         string `json:"path" yaml:"path"`
	RotTime      int64  `json:"rotTime" yaml:"rotTime"`
	RotSize      int64  `json:"rotSize" yaml:"rotSize"`
	RotCount     uint   `json:"rotCount" yaml:"rotCount"`
	MaxAge       int64  `json:"maxAge" yaml:"maxAge"`
	ForceNewFile bool   `json:"forceNewFile" yaml:"forceNewFile"`
}

func InitRot(cfg *RotConfig) (*rotatelogs.RotateLogs, error) {
	path, err := osext.ExecutableFolder()
	if err != nil {
		return nil, err
	}

	var options []rotatelogs.Option
	if cfg.RotTime > 0 {
		options = append(options, rotatelogs.WithRotationTime(time.Duration(cfg.RotTime)*time.Second))
	}

	if cfg.RotSize > 0 {
		options = append(options, rotatelogs.WithRotationSize(cfg.RotSize))
	}

	if cfg.RotCount > 0 {
		options = append(options, rotatelogs.WithRotationCount(cfg.RotCount))
	}

	if cfg.MaxAge > 0 {
		options = append(options, rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*time.Second))
	}

	if cfg.ForceNewFile {
		options = append(options, rotatelogs.ForceNewFile())
	}

	rotLog, err := rotatelogs.New(filepath.Join(path, cfg.Path), options...)
	if err != nil {
		return nil, err
	}

	return rotLog, err
}
