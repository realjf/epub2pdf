package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool = false
var MoreVerbose bool = false                  // output the converted output
var EbookConvertPath string = "ebook-convert" // the ebook-convert executable path, need to install calibre
var JobsNum int = 5
var OutputPath string = ""
var Recursive bool = false        // recursive directory
var ToBeConvertedPath string = "" // the directory to be converted
var DeleteSource bool = false     // delete the source file when convert successfully

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
