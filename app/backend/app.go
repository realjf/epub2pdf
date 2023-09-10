// #############################################################################
// # File: app.go                                                              #
// # Project: backend                                                          #
// # Created Date: 2023/09/11 00:19:54                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 00:22:32                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package backend

type BApp interface {
}

type bApp struct {
}

func NewBApp() BApp {
	b := &bApp{}

	return b
}

func (b *bApp) Run() {

}
