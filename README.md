# RotZap
Provider an easy way to initialize zap with file-rotatelogs.

# Installation
```go
go get github.com/ddzizz/rotzap
```

# Dependencies
- [github.com/uber-go/zap](https://github.com/uber-go/zap)
- [github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)

# Usage
## Init form yaml or json configuration file
```go
zapLog, err := rotzap.InitRotZapFromCfgFile("sample.yml")
if err != nil {
    fmt.Println(err.Error())
    return
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```

## Init from yaml configuration string
```go
zapLog, err := rotzap.InitRotZapFromYaml(yamlStr)
if err != nil {
    t.Fatal(err)
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```

## Init from json configuration string
```go
zapLog, err := rotzap.InitRotZapFromJSON(jsonStr)
if err != nil {
    t.Fatal(err)
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```