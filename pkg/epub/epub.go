// #############################################################################
// # File: epub.go                                                             #
// # Project: epub                                                             #
// # Created Date: 2024/12/09 22:22:17                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 22:52:14                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package epub

import "github.com/realjf/epub2pdf/pkg/epub/command"

type Invoker struct {
	commands []command.ICommand
}

func (i *Invoker) AddCommand(command command.ICommand) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}
