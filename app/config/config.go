// #############################################################################
// # File: config.go                                                           #
// # Project: config                                                           #
// # Created Date: 2023/09/10 23:15:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/10 23:31:29                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################

package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	flag.StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	flag.Parse()
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = os.Getenv("EPUB2PDF_CONFIG_PATH")
	}
	if cfgFile == "" {
		cfgFile = "./config.yaml"
	}

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("can not read config.yaml:", err)
		os.Exit(1)
	}
	viper.WatchConfig()
}
