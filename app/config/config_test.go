// #############################################################################
// # File: config_test.go                                                      #
// # Project: config                                                           #
// # Created Date: 2024/02/04 15:04:34                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/02/04 15:11:13                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package config_test

import (
	"fmt"
	"testing"

	"github.com/realjf/epub2pdf/app/config"
)

func TestInitConfig(t *testing.T) {
	config.InitConfigWithPath("../config.toml")
	fmt.Printf("%#v\n", config.GlobalConfig)
}
