package cmd

import "path/filepath"

const FILE_EXT_PDF = ".pdf"
const FILE_EXT_EPUB = ".epub"

type FileObj struct {
	name       string // just filename, without extension
	fileName   string // full filename
	ext        string // file extension
	abs        string // full filepath
	rootPath   string // root path
	toExt      string // convert format
	toRootPath string // output path
}

func NewFileObj(name, ext, rootpath, toext string) *FileObj {
	return &FileObj{
		name:       name,
		fileName:   name + ext,
		ext:        ext,
		abs:        filepath.Join(rootpath, name+ext),
		rootPath:   rootpath,
		toExt:      toext,
		toRootPath: "",
	}
}

func (f *FileObj) Name() string {
	return f.name
}

func (f *FileObj) FileName() string {
	return f.fileName
}

func (f *FileObj) Ext() string {
	return f.ext
}

func (f *FileObj) Abs() string {
	return f.abs
}

func (f *FileObj) RootPath() string {
	return f.rootPath
}

func (f *FileObj) ToExt() string {
	return f.toExt
}

func (f *FileObj) ToFileName() string {
	return f.name + f.toExt
}

func (f *FileObj) ToAbs() string {
	if f.toRootPath == "" {
		f.toRootPath = f.rootPath
	}
	return filepath.Join(f.toRootPath, f.name+f.toExt)
}

func (f *FileObj) ToRootPath(path string) *FileObj {
	if path == "" {
		f.toRootPath = f.rootPath
	} else {
		f.toRootPath = path
	}

	return f
}
