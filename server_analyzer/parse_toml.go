package server_analyzer

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config map[string]interface{}

// readConfig 读取配置
func readConfig(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Println("failed to decode config file:", err)
		return config, err
	}
	return config, nil
}

func isServerConfigOk(path string) error {
	var (
		config Config
		err    error
	)
	if config, err = readConfig(path); err != nil {
		return errors.New(fmt.Sprintf("配置文件解析失败: %v", path))
	}

	if config["server"] == nil {
		return errors.New(fmt.Sprintf("缺少服务配置"))
	}

	var (
		ok           bool
		serverConfig = make(map[string]interface{})
	)
	if serverConfig, ok = config["server"].(map[string]interface{}); !ok {
		return errors.New(fmt.Sprintf("服务配置异常1: %+v", config["server"]))
	}

	if serverConfig["port"] == "" || serverConfig["addr"] == "" {
		return errors.New(fmt.Sprintf("服务配置异常2: %+v", serverConfig))
	}

	return nil
}
