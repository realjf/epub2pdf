// #############################################################################
// # File: parse.go                                                            #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 16:43:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 22:26:06                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import "github.com/pkg/errors"

func (e *epub) parseFile() (err error) {
	opfFile, err := openOPF(e.zip)
	if err != nil {
		return errors.WithStack(err)
	}
	defer opfFile.Close()

	return
}
