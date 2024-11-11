// #############################################################################
// # File: config.go                                                           #
// # Project: config                                                           #
// # Created Date: 2023/09/10 23:15:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/11 11:31:35                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################

package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/realjf/zlog"
)

var cfgFile string
var GlobalConfig Config

func InitConfig() {
	if cfgFile == "" {
		cfgFile = os.Getenv("EPUB2PDF_CONFIG_PATH")
	}
	if cfgFile == "" {
		cfgFile = "./config.toml"
	}

	_, err := toml.DecodeFile(cfgFile, &GlobalConfig)
	if err != nil {
		panic(err)
	}
}

func InitConfigWithPath(path string) {
	cfgFile = path
	if cfgFile == "" {
		cfgFile = os.Getenv("EPUB2PDF_CONFIG_PATH")
	}
	if cfgFile == "" {
		cfgFile = "./config.toml"
	}

	_, err := toml.DecodeFile(cfgFile, &GlobalConfig)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Frontend FConfig `toml:"frontend"`
	Backend  BConfig `toml:"backend"`
}

type FConfig struct {
	Name    string `toml:"name"`
	ID      string `toml:"ID"`
	Icon    string `toml:"icon"`
	Version string `toml:"version"`
	Build   int64  `toml:"build"`
}

type BConfig struct {
	Log Log `toml:"log"`
}

type Log struct {
	Level    zlog.LogLevel `toml:"level"`    // 日志级别
	Compress bool          `toml:"compress"` // 日志是否压缩
	MaxAge   int           `toml:"maxAge"`   // 日志保留天数
	MaxSize  int           `toml:"maxSize"`  // 单个日志文件最大大小
	Filename string        `toml:"filename"` // 日志文件存储路径
}
