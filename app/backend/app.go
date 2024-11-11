// #############################################################################
// # File: app.go                                                              #
// # Project: backend                                                          #
// # Created Date: 2023/09/11 00:19:54                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/11 11:41:47                                        #
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

	"github.com/realjf/epub2pdf/app/backend/model"
	"github.com/realjf/epub2pdf/app/backend/utils"
	"github.com/realjf/zlog"
)

type BApp interface {
}

type bApp struct {
	lock sync.Mutex
	ctx  context.Context

	recursive bool

	inputPath  string
	outputPath string
}

func NewBApp() BApp {
	b := &bApp{
		lock: sync.Mutex{},
		ctx:  context.Background(),
	}

	return b
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
		zlog.ZLog().Error(err.Error())
		return files
	}

	if !fileinfo.IsDir() {
		// if it is a file
		fileObj := model.NewFileObj(utils.FileNameWithoutExtension(fileinfo.Name()), filepath.Ext(root), filepath.Dir(root), model.FILE_EXT_PDF)
		if !utils.FileExists(fileObj.Abs()) {
			zlog.ZLog().Warn("File[" + fileObj.Abs() + "] not found")
			return nil
		}
		files = append(files, fileObj)
		return files
	}

	// if it is a directory
	err = filepath.Walk(root,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				zlog.ZLog().Error(err.Error())
				return err
			}
			if root == ".." {
				return nil
			}

			var rootpath string
			if root == "." {
				rootpath, err = filepath.Abs(b.inputPath)
				if err != nil {
					zlog.ZLog().Error(err.Error())
					return err
				}
				rootpath = filepath.Join(rootpath, filepath.Dir(fp))
				zlog.ZLog().Debugf("Current directory1: %s", rootpath)
			} else {
				if info.IsDir() {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						zlog.ZLog().Error(err.Error())
						return err
					}
					zlog.ZLog().Debugf("Current directory2: %s", rootpath)

				} else {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						zlog.ZLog().Error(err.Error())
						return err
					}
					rootpath = filepath.Dir(rootpath)
					zlog.ZLog().Debugf("Current directory3: %s", rootpath)
				}

			}
			if !b.recursive {
				if ro, err := filepath.Abs(root); err != nil {
					zlog.ZLog().Error(err.Error())
					return err
				} else {
					if rootpath != ro {
						// Non recursive
						zlog.ZLog().Debug("Non recursive: " + rootpath + "," + ro)
						return nil
					}
				}
			}

			if !info.IsDir() && filepath.Ext(fp) != "" && formats[filepath.Ext(fp)] {
				fileObj := model.NewFileObj(utils.FileNameWithoutExtension(info.Name()), filepath.Ext(fp), rootpath, model.FILE_EXT_PDF)
				if !utils.FileExists(fileObj.Abs()) {
					zlog.ZLog().Warn("File[" + fileObj.Abs() + "] not found")
					return nil
				}
				files = append(files, fileObj)
				zlog.ZLog().Debug("The path[" + fileObj.Abs() + "] to be converted")
				return nil
			}

			return nil
		})
	if err != nil {
		zlog.ZLog().Fatal(err.Error())
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
			zlog.ZLog().Errorf("get output directory error: %s", err.Error())
			return
		}
		output_file = path.Join(abspath, file)
	}

	err := os.Rename(input_file, output_file)
	if err != nil {
		zlog.ZLog().Errorf("move %s error: %v", input_file, err)
	}
}
