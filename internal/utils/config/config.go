package config

import (
	"github.com/spf13/viper"
	"fmt"
)

const (
	// 配置文件扩展名
	fileExt = "yaml"
	// 配置文件名称
	fileName = "config"
)

func Init(path string)  {
	
	viper.SetConfigName(fileName) // 配置文件名称(无扩展名)
	viper.SetConfigType(fileExt) // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 默认路径
	viper.AddConfigPath("./config")

	// 如果自定义路径
	if path != "" {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig() // 查找并读取配置文件
	fmt.Println(err)


	// 监控配置文件变化
	viper.WatchConfig()
}