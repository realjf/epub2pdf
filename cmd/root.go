package cmd

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("can not read config.yaml:", err)
		os.Exit(1)
	}
	if logLevel := viper.GetBool("log.debug"); logLevel {
		log.SetLevel(log.DebugLevel)
	}

	// check ebook-convert exist
	EbookConvertPath = viper.GetString("ebookconvert.path")
	if _, err := exec.LookPath(EbookConvertPath); err != nil {
		log.Fatal("ebook-convert is not in your PAHT")
		os.Exit(1)
	}
}
