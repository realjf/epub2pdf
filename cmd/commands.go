package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "convert epub to pdf",
	Long:  `Convert the specified directory epub file to pdf file`,
	Run: func(cmd *cobra.Command, args []string) {
		// clean output directory
		CleanDir()
		//
		if len(args) > 1 {
			if runtime.GOOS == "windows" {
				panic("Usage: epub2pdf.exe directory")
			} else if runtime.GOOS == "linux" {
				panic("Usage: ./epub2pdf directory")
			} else if runtime.GOOS == "darwin" {
				panic("Usage: ./epub2pdf directory")
			} else {
				panic("Usage: ./epub2pdf directory")
			}

		}

		// start to convert
		Convert(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// clean output directory
func CleanDir() {
	dir := "./output"
	d, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	files, err := d.Readdirnames(-1)
	if err != nil {
		panic(err)
	}
	for _, name := range files {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			panic(err)
		}
	}
	log.Info("output directory is clean")
}
