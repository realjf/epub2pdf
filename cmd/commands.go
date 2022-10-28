package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool
var EbookConvertPath string = ""
var JobsNum int

var rootCmd = &cobra.Command{
	Use:   "epub2pdf <EPUB-DIRECTORY>",
	Short: "convert epub to pdf",
	Long:  `Convert the specified directory epub file to pdf file`,
	Run: func(cmd *cobra.Command, args []string) {
		// clean output directory
		CleanDir()
		//
		if len(args) > 1 {
			if runtime.GOOS == "windows" {
				panic("Usage: epub2pdf.exe <directory>")
			} else if runtime.GOOS == "linux" {
				panic("Usage: ./epub2pdf <directory>")
			} else if runtime.GOOS == "darwin" {
				panic("Usage: ./epub2pdf <directory>")
			} else {
				panic("Usage: ./epub2pdf <directory>")
			}

		}

		// start to convert
		Convert(args[0])
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.Flags().StringVarP(&EbookConvertPath, "ebook-convert-path", "p", "ebook-convert", "The ebook-convert path")
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	rootCmd.Flags().IntVarP(&JobsNum, "jobs", "j", 5, "Allow N jobs at once; infinite jobs with no arg")
	// check ebook-convert exist
	if _, err := exec.LookPath(EbookConvertPath); err != nil {
		log.Fatal("ebook-convert is not in your PAHT")
		os.Exit(1)
	}
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}
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
	err := makeDirectoryIfNotExists(dir)
	if err != nil {
		panic(err)
	}
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

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
