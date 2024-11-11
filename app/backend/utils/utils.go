// #############################################################################
// # File: utils.go                                                            #
// # Project: utils                                                            #
// # Created Date: 2023/09/11 07:41:28                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/11 11:42:33                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/realjf/zlog"
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
func CleanDir(dir string) {
	err := MakeDirectoryIfNotExists(dir)
	if err != nil {
		zlog.ZLog().Error(err.Error())
		return
	}
	d, err := os.Open(dir)
	if err != nil {
		zlog.ZLog().Error(err.Error())
		return
	}
	defer d.Close()

	files, err := d.Readdirnames(-1)
	if err != nil {
		zlog.ZLog().Error(err.Error())
		return
	}
	for _, name := range files {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			zlog.ZLog().Error(err.Error())
			return
		}
	}
	zlog.ZLog().Info("output directory is clean")
}

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func DeleteFile(file string) error {
	return os.Remove(file)
}
