package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
}

func initConfig() {
	// if cfgFile != "" {
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	viper.SetConfigFile("./config.yaml")
	// }

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("can not read config.yaml:", err)
	// 	os.Exit(1)
	// }
}
