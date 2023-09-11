// #############################################################################
// # File: request.go                                                          #
// # Project: model                                                            #
// # Created Date: 2023/09/11 07:56:56                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 08:07:07                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package model

import "time"

type EpubToPDFReq struct {
	InputPath  string
	OutputPath string
	JobsNum    int
	Timeout    time.Duration
	IsDelete   bool // Delete source file after converted done
}

type EpubToPDFRes struct {
}
