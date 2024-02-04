// #############################################################################
// # File: config.go                                                           #
// # Project: config                                                           #
// # Created Date: 2023/09/10 23:15:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/02/04 15:16:18                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################

package config

import (
	"os"

	"github.com/BurntSushi/toml"
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
	Level string `toml:"level"`
}
