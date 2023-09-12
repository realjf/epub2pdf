// #############################################################################
// # File: epub_load_test.go                                                   #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 17:13:13                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 21:32:19                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"testing"
)

func TestOpenEpub(t *testing.T) {
	e, err := OpenEpub("/home/realjf/Downloads/a78fa0b1-ee40-460f-9c0a-3a4ffab81287.epub")
	if err != nil {
		t.Fatalf("%+v\n", err)
		return
	}
	defer e.Close()
	return
}
