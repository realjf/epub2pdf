// #############################################################################
// # File: app.go                                                              #
// # Project: backend                                                          #
// # Created Date: 2023/09/11 00:19:54                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 15:40:24                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package backend

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/realjf/epub2pdf/app/backend/model"
	"github.com/realjf/epub2pdf/app/backend/utils"
)

type BApp interface {
}

type bApp struct {
	logLevel string
	lock     sync.Mutex
	ctx      context.Context

	recursive bool

	inputPath  string
	outputPath string
}

func NewBApp() BApp {
	b := &bApp{
		lock:     sync.Mutex{},
		logLevel: "debug",
		ctx:      context.Background(),
	}

	lvl, err := log.ParseLevel(b.logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(lvl)

	return b
}

func (b *bApp) SetLogLvel(level string) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Error(err)
		return err
	}
	b.logLevel = level
	log.SetLevel(lvl)

	return nil
}

func (b *bApp) SetRecursive(recursive bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.recursive = recursive
}

func (b *bApp) SetInputPath(inputPath string) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.inputPath = inputPath
}

func (b *bApp) SetOutputPath(outputPath string) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.outputPath = outputPath
}

func (b *bApp) Run() {

}

func (b *bApp) getPaths() []*model.FileObj {
	b.lock.Lock()
	defer b.lock.Unlock()
	root := b.inputPath
	formats := map[string]bool{model.FILE_EXT_EPUB: true}

	files := []*model.FileObj{}

	fileinfo, err := os.Stat(root)
	if os.IsNotExist(err) {
		log.Error(err)
		return files
	}

	if !fileinfo.IsDir() {
		// if it is a file
		fileObj := model.NewFileObj(utils.FileNameWithoutExtension(fileinfo.Name()), filepath.Ext(root), filepath.Dir(root), model.FILE_EXT_PDF)
		if !utils.FileExists(fileObj.Abs()) {
			log.Warn("File[" + fileObj.Abs() + "] not found")
			return nil
		}
		files = append(files, fileObj)
		return files
	}

	// if it is a directory
	err = filepath.Walk(root,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				log.Error(err)
				return err
			}
			if root == ".." {
				return nil
			}

			var rootpath string
			if root == "." {
				rootpath, err = filepath.Abs(b.inputPath)
				if err != nil {
					log.Error(err)
					return err
				}
				rootpath = filepath.Join(rootpath, filepath.Dir(fp))
				log.Trace("Current directory1: " + rootpath)
			} else {
				if info.IsDir() {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						log.Error(err)
						return err
					}
					log.Trace("Current directory2: " + rootpath)

				} else {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						log.Error(err)
						return err
					}
					rootpath = filepath.Dir(rootpath)
					log.Trace("Current directory3: " + rootpath)
				}

			}
			if !b.recursive {
				if ro, err := filepath.Abs(root); err != nil {
					log.Error(err)
					return err
				} else {
					if rootpath != ro {
						// Non recursive
						log.Debug("Non recursive: " + rootpath + "," + ro)
						return nil
					}
				}
			}

			if !info.IsDir() && filepath.Ext(fp) != "" && formats[filepath.Ext(fp)] {
				fileObj := model.NewFileObj(utils.FileNameWithoutExtension(info.Name()), filepath.Ext(fp), rootpath, model.FILE_EXT_PDF)
				if !utils.FileExists(fileObj.Abs()) {
					log.Warn("File[" + fileObj.Abs() + "] not found")
					return nil
				}
				files = append(files, fileObj)
				log.Debug("The path[" + fileObj.Abs() + "] to be converted")
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

func (b *bApp) moveToOutput(rootpath, file string) {
	input_file := path.Join(rootpath, file)
	output_file := path.Join(rootpath, file)
	if b.outputPath != "" {
		abspath, err := filepath.Abs(b.outputPath)
		if err != nil {
			log.Errorf("get output directory error: %s", err.Error())
			return
		}
		output_file = path.Join(abspath, file)
	}

	err := os.Rename(input_file, output_file)
	if err != nil {
		log.Errorf("move %s error: %v", input_file, err)
	}
}
