package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool
var EbookConvertPath string = "ebook-convert" // the ebook-convert executable path, need to install calibre
var JobsNum int
var OutputPath string = ""
var Recursive bool                // recursive directory
var ToBeConvertedPath string = "" // the directory to be converted

var rootCmd = &cobra.Command{
	Use: "epub2pdf",
	Run: func(cmd *cobra.Command, args []string) {

	},
	Version: CurrentVersion,
	Args:    cobra.MinimumNArgs(1),
}

func Execute() {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
