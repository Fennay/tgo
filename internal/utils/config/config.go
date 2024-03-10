package config

import (
	"github.com/spf13/viper"
)

const (
	// 配置文件扩展名
	fileExt = "yaml"
	// 配置文件名称
	fileName = "config"
)

// Config 配置项
type Config struct {
	Http  Http
	DB    DB
	Redis Redis
	Log   Log
}

// Http 配置
type Http struct {
	URL     string
	Port    int
	RunMode string
}

// DB 数据库
type DB struct {
	Dialect  string
	URL      string
	Port     int
	Database string
	User     string
	Password string
}

// Redis redis配置
type Redis struct {
	URL      string
	Port     int
	DB       int
	Password string
}

// Log 日志
type Log struct {
	Output string
	Level  string
}

var configStruct *Config

// Init 初始化配置
func Init(path ...string) {
	newPath := resolvePath(path)
	viper.SetConfigName(fileName) // 配置文件名称(无扩展名)
	viper.SetConfigType(fileExt)  // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 默认路径
	viper.AddConfigPath("./config")

	// 如果自定义路径
	if newPath != "" {
		viper.AddConfigPath(newPath)
	}

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(err)
	}

	// 序列化
	err = viper.Unmarshal(&configStruct)
	if err != nil {
		panic(err)
	}

	// 设置默认值
	if configStruct.Http.URL == "" {
		configStruct.Http.URL = "127.0.0.1"
	}
	if configStruct.Http.Port == 0 {
		configStruct.Http.Port = 8088
	}

	// 监控配置文件变化
	viper.WatchConfig()
}

func resolvePath(path []string) string {
	if len(path) > 0 {
		return path[0]
	}
	// 如果没传递使用默认配置文件
	return ""
}

// GetConfig 获得配置
func GetConfig() *Config {
	if configStruct != nil {
		return configStruct
	}

	Init("")
	return configStruct
}
