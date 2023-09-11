// #############################################################################
// # File: req_and_res.go                                                      #
// # Project: model                                                            #
// # Created Date: 2023/09/11 12:52:00                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 15:47:36                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package model

import (
	"time"
)

type ConvertReq struct {
	InputFiles []*FileObj
	OutputPath string
	JobsNum    int
	Timeout    time.Duration
	IsDelete   bool // Delete source file after converted done
}

type ConvertRes struct {
}
