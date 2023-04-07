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
type ConfigS struct {
	Http  Http
	DB    DB
	Redis Redis
	Log   Log
}

// Http 配置
type Http struct {
	Port    int
	Runmode string
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

// Redis redis
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

var configData *ConfigS

// Init 初始化配置
func Init(path string) {

	viper.SetConfigName(fileName) // 配置文件名称(无扩展名)
	viper.SetConfigType(fileExt)  // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 默认路径
	viper.AddConfigPath("./config")

	// 如果自定义路径
	if path != "" {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(err)
	}

	//序列化
	err = viper.Unmarshal(&configData)
	if err != nil {
		panic(err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
}

// GetConfig 获得配置
func GetConfig() *ConfigS {
	if configData != nil {
		return configData
	}

	Init("")
	return configData
}
