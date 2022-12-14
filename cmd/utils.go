package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
	log "github.com/sirupsen/logrus"
)

// return true means file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// return filename without extension
func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
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

func moveToOutput(rootpath, file string) {
	input_file := path.Join(rootpath, file)
	output_file := path.Join(rootpath, file)
	if OutputPath != "" {
		abspath, err := filepath.Abs(OutputPath)
		if err != nil {
			log.Error("get output directory error: %s", err.Error())
			return
		}
		output_file = path.Join(abspath, file)
	}

	err := os.Rename(input_file, output_file)
	if err != nil {
		log.Error(color.InRed("move " + input_file + " error: " + err.Error()))
	}
}

func deleteFile(file string) error {
	return os.Remove(file)
}
