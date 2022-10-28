package main

import (
	"epub2pdf/cmd"
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
