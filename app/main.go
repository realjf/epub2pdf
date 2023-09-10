// #############################################################################
// # File: main.go                                                             #
// # Project: app                                                              #
// # Created Date: 2023/09/10 23:19:37                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 00:26:10                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	_ "github.com/realjf/epub2pdf/app/config"
	"github.com/realjf/epub2pdf/app/frontend"
)

var Version string = ""

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	title := viper.GetString("frontend.title")
	fapp := frontend.NewApp(title)
	fapp.Run()

	for {
		select {
		case <-quit:
			return
		}
	}
}
