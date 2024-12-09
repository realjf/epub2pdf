// #############################################################################
// # File: extract.go                                                          #
// # Project: extract                                                          #
// # Created Date: 2024/12/09 22:25:20                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 22:56:44                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package extract

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ExtractCommand struct {
	epubFilePath string
	cacheDir     string
	reader       *zip.ReadCloser
}

func New(epubFilePath string, cacheDir string) *ExtractCommand {
	return &ExtractCommand{
		epubFilePath: epubFilePath,
		cacheDir:     cacheDir,
	}
}

func (e *ExtractCommand) Execute() (err error) {
	// open epub file
	e.reader, err = zip.OpenReader(e.epubFilePath)
	if err != nil {
		fmt.Printf("[ExtractCommand] open epub file error: %v\n", err)
		return
	}

	// make cache directory
	if err := os.MkdirAll(e.cacheDir, os.ModePerm); err != nil {
		fmt.Printf("[ExtractCommand] make cache directory error: %v\n", err)
		return err
	}

	for _, file := range e.reader.File {
		err = e.extract(file)
		if err != nil {
			fmt.Printf("[ExtractCommand] extract zip file error: %v\n", err)
			return
		}
	}

	return
}

func (e *ExtractCommand) Close() {
	e.reader.Close()
}

func (e *ExtractCommand) extract(file *zip.File) (err error) {
	srcFile, err := file.Open()
	if err != nil {
		fmt.Printf("[ExtractCommand] open zip file error: %v\n", err)
		return
	}
	defer srcFile.Close()

	destPath := filepath.Join(e.cacheDir, file.Name)

	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
			fmt.Printf("[ExtractCommand] make cache directory[%s] error: %v\n", destPath, err)
			return err
		}
		return nil
	}

	// make sure parent directory exist
	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		fmt.Printf("[ExtractCommand] make cache root directory error: %v\n", err)
		return err
	}

	// create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("[ExtractCommand] create cache file error: %v\n", err)
		return
	}
	defer destFile.Close()

	// copy content to the cache file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Printf("[ExtractCommand] copy content to cache file error: %v\n", err)
		return
	}
	return
}
