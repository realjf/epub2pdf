// #############################################################################
// # File: meta_test.go                                                        #
// # Project: meta                                                             #
// # Created Date: 2024/12/09 23:00:48                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 23:27:31                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package meta_test

import (
	"testing"

	"github.com/realjf/epub2pdf/pkg/epub/meta"
)

func TestMetaParse(t *testing.T) {
	m := meta.New("./.cache", "")
	defer m.Close()
	m.Execute()
}
