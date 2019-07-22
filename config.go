package util

import (
// "log"
)

// ReadConfig 读取配置文件
// 路径为空时加载 "./config/config_default.json" 文件
func ReadConfig(filename string) ([]byte, *ErrorStruct) {
	if filename == "" {
		filename = "./config/config_default.json"
	}

	bytes, err := ReadeFile(filename)
	if err != nil {
		return nil, RetErr(400, err.Error())
	}
	return bytes, nil
}
