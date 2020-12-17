package rotzap

import (
	"path/filepath"
	"testing"

	"github.com/kardianos/osext"
	"go.uber.org/zap"
)

var yamlStr = `
rot:
  path: "log/yaml/%Y%m%d_%H%M%S"
  rotTime: 60
  rotSize: 10240

zap:
  development: true
  disableCaller: false
  disableStacktrace: false
  encoderConfig:
    timeKey: "ts"
    levelKey: "lev"
    messageKey: "msg"
    
`

var jsonStr = `
{
    "rot": {
        "path": "log/json/%Y%m%d_%H%M%S",
        "rotTime": 60,
        "rotSize": 10240
    },
    "zap": {
        "development": true,
        "disableCaller": false,
        "disableStacktrace": false,
		"encoding": "json",
        "encoderConfig": {
            "timeKey": "ts",
            "levelKey": "lev",
            "messageKey": "msg"
        }
    }
}
`

func TestInitRotZapFromYaml(t *testing.T) {
	zapLog, err := InitRotZapFromYaml(yamlStr)
	if err != nil {
		t.Fatal(err)
	}
	defer zapLog.Sync()

	zapLog.Info("[TestInitRotZapFromYaml] RotZap provide an easy way to initialize zap with file-rotatelogs")
	zapLog.Error("[TestInitRotZapFromYaml] RotZap provide an easy way to initialize zap with file-rotatelogs")
}

func TestInitRotZapFromJson(t *testing.T) {
	zapLog, err := InitRotZapFromJSON(jsonStr)
	if err != nil {
		t.Fatal(err)
	}
	defer zapLog.Sync()

	zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs", zap.String("test", "TestInitRotZapFromJson"))
	zapLog.Error("RotZap provide an easy way to initialize zap with file-rotatelogs", zap.String("test", "TestInitRotZapFromJson"))
}

func TestInitRotZapFromCfgFile(t *testing.T) {
	exePath, err := osext.ExecutableFolder()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(exePath, "sample")

	zapLog, err := InitRotZapFromCfgFile(path)
	if err != nil {
		t.Fatal(err)
	}
	defer zapLog.Sync()

	zapLog.Info("[TestInitRotZapFromCfgFile] RotZap provide an easy way to initialize zap with file-rotatelogs")
	zapLog.Error("[TestInitRotZapFromCfgFile] RotZap provide an easy way to initialize zap with file-rotatelogs")
}
