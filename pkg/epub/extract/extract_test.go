// #############################################################################
// # File: extract_test.go                                                     #
// # Project: extract                                                          #
// # Created Date: 2024/12/09 22:45:27                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 22:50:27                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package extract_test

import (
	"testing"

	"github.com/realjf/epub2pdf/pkg/epub/extract"
)

func TestExtract(t *testing.T) {
	file := "./60c28a4a-a255-42d1-a045-3a82da3da968.epub"
	cacheDir := "./.cache"
	cmd := extract.New(file, cacheDir)
	defer cmd.Close()
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("done")
}
