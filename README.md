# rotzap
Provider an easy way to initialize zap with file-rotatelogs.

# Usage
```go
zapLog, err := InitRotZapFromCfgFile("D:/Projects/ddzizz/rotzap/sample.yml")
if err != nil {
    fmt.Println(err.Error())
    return
}
defer zapLog.Sync()

zapLog.Info("ZapLog provide an easy way to initialize zap with file-rotatelogs")
```