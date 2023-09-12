// #############################################################################
// # File: epub_load_test.go                                                   #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 17:13:13                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 17:24:07                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import "testing"

func TestOpenEpub(t *testing.T) {
	e, err := OpenEpub("~/Downloads/8d3368bd-baf3-4cf6-b1d6-08358193b6b1.epub")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer e.Close()

}
