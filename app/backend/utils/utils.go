// #############################################################################
// # File: utils.go                                                            #
// # Project: utils                                                            #
// # Created Date: 2023/09/11 07:41:28                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 07:51:31                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package utils

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
	log "github.com/sirupsen/logrus"

	"github.com/realjf/epub2pdf/app/backend/model"
)

// return true means file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// return filename without extension
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}

// clean output directory
func CleanDir() {
	dir := "./output"
	err := MakeDirectoryIfNotExists(dir)
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

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func MoveToOutput(rootpath, file string) {
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

func DeleteFile(file string) error {
	return os.Remove(file)
}

func GetPaths(root string) []*model.FileObj {
	formats := map[string]bool{model.FILE_EXT_EPUB: true}

	files := []*model.FileObj{}
	err := filepath.Walk(root,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if root == ".." {
				return nil
			}

			var rootpath string
			if root == "." {
				rootpath, err = filepath.Abs(ToBeConvertedPath)
				if err != nil {
					return err
				}
				rootpath = filepath.Join(rootpath, filepath.Dir(fp))
				if MoreVerbose {
					log.Info("Current directory1: " + rootpath)
				}
			} else {
				if info.IsDir() {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						return err
					}
					if MoreVerbose {
						log.Info("Current directory2: " + rootpath)
					}
				} else {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						return err
					}
					rootpath = filepath.Dir(rootpath)
					if MoreVerbose {
						log.Info("Current directory3: " + rootpath)
					}
				}

			}
			if !Recursive {
				if ro, err := filepath.Abs(root); err != nil {
					return err
				} else {
					if rootpath != ro {
						// Non recursive
						if MoreVerbose {
							log.Info("Non recursive: " + rootpath + "," + ro)
						}
						return nil
					}
				}
			}

			if !info.IsDir() && filepath.Ext(fp) != "" && formats[filepath.Ext(fp)] {
				fileObj := model.NewFileObj(FileNameWithoutExtension(info.Name()), filepath.Ext(fp), rootpath, model.FILE_EXT_PDF)
				if !FileExists(fileObj.Abs()) {
					log.Warn(color.InYellow("File[" + fileObj.Abs() + "] not found"))
					return nil
				}
				files = append(files, fileObj)
				if Verbose {
					log.Info("The path[" + fileObj.Abs() + "] to be converted")
				}
				return nil
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	// return strings
	return files

}
