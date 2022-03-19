package conf

import (
	"fmt"
	"path"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// LogLevel 加载配置输出日志级别
type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelCritical
	LevelFatal
)

var SupportedExtensions = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl"}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SetLogLevel(level LogLevel) {
	// 打印 viper 配置文件加载过程
	jwwLogLevel := jww.Threshold(level)
	jww.SetStdoutThreshold(jwwLogLevel)
}

func LoadConfig(confPath string, rawVal interface{}, logLevel LogLevel) (err error) {
	confPath = strings.Replace(confPath, "\\", "/", -1)
	fileDir := path.Dir(confPath)
	fileFullName := path.Base(confPath)
	fileExtension := path.Ext(fileFullName)
	if fileExtension == "" {
		fileExtension = ".yaml"
	}
	fileType := fileExtension[1:]
	if !stringInSlice(fileType, SupportedExtensions) {
		return viper.UnsupportedConfigError(fileType)
	}
	nameLen := len(fileFullName) - len(fileExtension)
	configName := fileFullName[:nameLen]

	SetLogLevel(logLevel)
	conf := viper.New()
	conf.SetConfigName(configName) // 配置文件的名字
	conf.SetConfigType(fileType)   // 配置文件的类型
	conf.AddConfigPath(fileDir)    // 配置文件的路径

	if err = conf.ReadInConfig(); err != nil {
		return err
	}
	if err = conf.Unmarshal(rawVal); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
	return nil
}
