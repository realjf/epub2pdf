// #############################################################################
// # File: main.go                                                             #
// # Project: app                                                              #
// # Created Date: 2023/09/10 23:19:37                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/02/04 15:36:48                                        #
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
)

var Version string = ""

func main() {
	config.InitConfig()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	appName := config.GlobalConfig.Frontend.Name
	fapp := frontend.NewApp(appName)
	fapp.Run()

	for {
		select {
		case <-quit:
			return
		}
	}
}
