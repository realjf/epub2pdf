// #############################################################################
// # File: epub_new.go                                                         #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 16:37:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 16:54:19                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"net/http"

	"github.com/gofrs/uuid"
)

func NewEpub(title string) Epub {
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
		toc: newToc(),
	}

	e.SetIdentifier(urnUUIDPrefix + uuid.Must(uuid.NewV4()).String())
	e.SetLang(defaultEpubLang)
	if title != "" {
		e.SetTitle(title)
	}
	return e
}
