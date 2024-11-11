// #############################################################################
// # File: main.go                                                             #
// # Project: app                                                              #
// # Created Date: 2023/09/10 23:19:37                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/11 13:04:58                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/realjf/epub2pdf/app/config"
	_ "github.com/realjf/epub2pdf/app/config"
	"github.com/realjf/epub2pdf/app/frontend"
	"github.com/realjf/zlog"
)

var Version string = ""

func main() {
	config.InitConfig()
	zlog.InitZLog(&zlog.ZLogConfig{
		Level:    config.GlobalConfig.Backend.Log.Level,
		Compress: config.GlobalConfig.Backend.Log.Compress,
		LogMode:  "file",
		Encoding: "json",
		MaxSize:  config.GlobalConfig.Backend.Log.MaxSize,
		MaxAge:   config.GlobalConfig.Backend.Log.MaxAge,
		LogFile:  config.GlobalConfig.Backend.Log.Filename,
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	appName := config.GlobalConfig.Frontend.Name
	fapp := frontend.NewApp(appName)
	fapp.Run()

	for {
		select {
		case <-fapp.IsShutdown():
			zlog.ZLog().Info("shutdown gracefully")
			return
		case <-quit:
			zlog.ZLog().Info("shutdown ungracefully")
			return
		}
	}
}
