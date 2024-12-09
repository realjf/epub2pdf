// #############################################################################
// # File: command.go                                                          #
// # Project: command                                                          #
// # Created Date: 2024/12/09 22:24:46                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 22:30:01                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package command

type ICommand interface {
	Execute() (err error)
}
