# RotZap
Provider an easy way to initialize zap with file-rotatelogs.

# Installation
```go
go get github.com/ddzizz/rotzap
```

# Dependencies
- [github.com/uber-go/zap](https://github.com/uber-go/zap)
- [github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
- [github.com/kardianos/osext](https://github.com/kardianos/osext)

# Usage
## Init from yaml or json configuration file
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

# Configurations
## Configurations for file-rotatelogs
For details, please check options.go in [github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs).
| Field name | Value type | Description |
| ---- | ---- | ---- |
| path | string | log filename pattern. |
| rotTime | int64 | sets the time between rotation. |
| rotSize | int64 | sets the log file size between rotation. |
| rotCount | uint | sets the number of files should be kept before it gets purged from the file system. |
| maxAge | int64 | sets the max age of a log file before it gets purged from  the file system. |
| forceNewFile | bool | ensures a new file is created every time New() is called. If the base file name already exists, an implicit rotation is performed |

## Configurations for zap
For details, please check config.go in [github.com/uber-go/zap](https://github.com/uber-go/zap).
