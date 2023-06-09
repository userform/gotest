package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	// "github.com/caarlos0/env/v6"
)

// 如何把配置映射成Config对象

// 从配置文件中加载配置
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}
	return nil
}

// 从环境遍历中加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}
