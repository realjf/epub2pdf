// #############################################################################
// # File: reader.go                                                           #
// # Project: reader                                                           #
// # Created Date: 2024/12/09 22:53:10                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 23:09:57                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package reader

import (
	"fmt"
	"os"
)

type ReaderCommand struct {
	filePath string
	content  []byte
}

func New(filePath string) *ReaderCommand {
	return &ReaderCommand{
		filePath: filePath,
	}
}

func (r *ReaderCommand) Execute() (err error) {
	r.content, err = os.ReadFile(r.filePath)
	if err != nil {
		fmt.Printf("[ReaderCommand] read from %s error: %v", r.filePath, err)
		return
	}
	return
}

func (r *ReaderCommand) Close() {

}

func (r *ReaderCommand) GetContent() []byte {
	return r.content
}
