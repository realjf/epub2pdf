package main

import (
	"epub2pdf/cmd"
	"runtime"
)

var Version string = ""

func main() {
	cmd.CurrentVersion = Version
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
