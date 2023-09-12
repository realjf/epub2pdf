// #############################################################################
// # File: epub_load.go                                                        #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 16:37:30                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 17:10:26                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
)

func OpenEpub(filename string) (e Epub, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	e1 := &epub{
		Client: http.DefaultClient,
		css:    make(map[string]string),
		fonts:  make(map[string]string),
		images: make(map[string]string),
		videos: make(map[string]string),
		audios: make(map[string]string),
		cover: &epubCover{
			cssFilename:   "",
			cssTempFile:   "",
			imageFilename: "",
			xhtmlFilename: "",
		},
		toc:  newToc(),
		file: file,
	}

	err = e1.load(e1.file, fileInfo.Size())
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *epub) load(r io.ReaderAt, size int64) (err error) {
	e.zip, err = zip.NewReader(r, size)
	if err != nil {
		return
	}

	return e.parseFile()
}

