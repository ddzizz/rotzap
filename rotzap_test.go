package rotzap

import (
	"testing"
)

func TestRotZapAll(t *testing.T) {
	zapLog, err := InitRotZapFromCfgFile("D:/Projects/ddzizz/rotzap/sample.yml")
	if err != nil {
		t.Fatal(err)
	}
	defer zapLog.Sync()

	zapLog.Info("ZapLog provide an easy way to initialize zap with file-rotatelogs")
	zapLog.Error("ZapLog provide an easy way to initialize zap with file-rotatelogs")
}
