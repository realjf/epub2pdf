package main

import (
	"runtime"

	"github.com/realjf/epub2pdf/cmd"
)

var Version string = ""

func main() {
	cmd.CurrentVersion = Version
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
