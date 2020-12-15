# RotZap
Provider an easy way to initialize zap with file-rotatelogs.

# Install
```go
go get github.com/ddzizz/rotzap
```

# Usage
## Init form yaml or json configuration file
```go
zapLog, err := InitRotZapFromCfgFile("sample.yml")
if err != nil {
    fmt.Println(err.Error())
    return
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```

## Init from yaml configuration string
```go
zapLog, err := InitRotZapFromYaml(yamlStr)
if err != nil {
    t.Fatal(err)
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```

## Init from json configuration string
```go
zapLog, err := InitRotZapFromJSON(jsonStr)
if err != nil {
    t.Fatal(err)
}
defer zapLog.Sync()

zapLog.Info("RotZap provide an easy way to initialize zap with file-rotatelogs")
```