// #############################################################################
// # File: epub_load.go                                                        #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 16:37:30                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 21:50:17                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

func OpenEpub(filename string) (Epub, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	e := &epub{
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

	err = e.load(e.file, fileInfo.Size())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return e, nil
}

// ============================================ load epub file ==============================================

func (e *epub) load(r io.ReaderAt, size int64) (err error) {
	e.zip, err = zip.NewReader(r, size)
	if err != nil {
		return
	}

	e.rootPath, err = getRootPath(e.zip)
	if err != nil {
		return errors.WithStack(err)
	}

	return e.parseFile()
}

type containerXML struct {
	// FIXME: only support for one rootfile, can it be more than one?
	Rootfile rootfile `xml:"rootfiles>rootfile"`
}
type rootfile struct {
	Path string `xml:"full-path,attr"`
}

func openOPF(file *zip.Reader) (io.ReadCloser, error) {
	path, err := getOpfPath(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return openFile(file, path)
}

func getRootPath(file *zip.Reader) (string, error) {
	opfPath, err := getOpfPath(file)
	if err != nil {
		return "", errors.WithStack(err)
	}
	pathDir := path.Dir(opfPath)
	if pathDir == "." {
		return "", nil
	} else {
		return path.Dir(opfPath) + "/", nil
	}
}

func getOpfPath(file *zip.Reader) (string, error) {
	f, err := openFile(file, "META-INF/container.xml")
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer f.Close()

	var c containerXML
	err = decodeXML(f, &c)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return c.Rootfile.Path, nil
}

func decodeXML(file io.Reader, v interface{}) error {
	decoder := xml.NewDecoder(file)
	decoder.Entity = xml.HTMLEntity
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(v)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func openFile(file *zip.Reader, path string) (io.ReadCloser, error) {
	for _, f := range file.File {
		if f.Name == path {
			reader, err := f.Open()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return reader, nil
		}
	}

	pathLower := strings.ToLower(path)
	for _, f := range file.File {
		if strings.ToLower(f.Name) == pathLower {
			reader, err := f.Open()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return reader, nil
		}
	}

	return nil, errors.WithStack(errors.New("File " + path + " not found"))
}
