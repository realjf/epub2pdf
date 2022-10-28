package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/TwiN/go-color"
	"github.com/realjf/gopool"
	log "github.com/sirupsen/logrus"
)

func moveToOutput(rootpath, file string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Error("get current directory error: %s", err.Error())
		return
	}
	input_file := path.Join(rootpath, file)
	output_file := path.Join(dir, "/output/"+file)
	err = os.Rename(input_file, output_file)
	if err != nil {
		log.Error(color.InRed("move " + input_file + " error: " + err.Error()))
	}
}

func deleteEpub(root, file string) {
	err := os.Remove(root + file)
	if err != nil {
		log.Fatal(err)
	}
}

func getPaths(root string) []string {
	var paths []string
	formats := map[string]bool{"epub": true}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".") {
			continue
		}

		f := strings.Split(file.Name(), ".")
		fname := f[0]
		ex := f[1]
		formatOk := formats[ex]
		if !formatOk {
			// filter epub format
			continue
		}
		// add filename to slice
		paths = append(paths, fname)
	}

	// return strings
	return paths

}

func Convert(rootpath string) {
	convertPool := gopool.NewPool(10)
	files := getPaths(rootpath)
	convertPool.SetTaskNum(len(files))

	// add task
	go func() {
		for _, filename := range files {
			x := filename
			myTaskFunc := func(param interface{}) (err error, r interface{}) {
				r, ok := param.(string)
				if !ok {
					return errors.New("task parameter is not string type"), r
				}
				input_file := path.Join(rootpath, x+".epub")
				output_file := path.Join(rootpath, x+".pdf")
				log.Infof("ready to convert %s to %s ...\n", input_file, output_file)

				cmd := exec.Command("ebook-convert", input_file, output_file)
				cmd.Env = os.Environ()
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					log.Error("Failed creating command stdoutpipe: ", err)
					return err, r
				}
				defer stdout.Close()
				stdoutReader := bufio.NewReader(stdout)

				stderr, err := cmd.StderrPipe()
				if err != nil {
					log.Error("Failed creating command stderrpipe: ", err)
					return err, r
				}
				defer stderr.Close()
				stderrReader := bufio.NewReader(stderr)

				if err = cmd.Start(); err != nil {
					log.Error("Failed starting command: ", err)
					return err, r
				}

				go handleReader(stdoutReader)
				go handleReader(stderrReader)

				if err := cmd.Wait(); err != nil {
					if exiterr, ok := err.(*exec.ExitError); ok {
						if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
							log.Debug("Exit Status: ", status.ExitStatus())
							return err, r
						}
					}
					return err, r
				}

				if err != nil {
					log.Error(color.InRed("======== failed to convert " + input_file + " =========="))
					log.Error(color.InRed(err.Error()))
					return
				}

				return
			}
			myTaskCallbackFunc := func(param interface{}) (e error, r interface{}) {
				input_file := path.Join(rootpath, param.(string)+".epub")
				log.Info(color.InGreen("======== convert " + input_file + " successfully =========="))
				return
			}
			task := gopool.NewTask(myTaskFunc, myTaskCallbackFunc, x)
			err := convertPool.AddTask(task)
			if err != nil {
				panic("add task error")
			}

		}
	}()

	convertPool.Run()
	log.Info(color.InGreen("tasks is completed!!!"))

	log.Info("ready to move pdf files...")
	for _, x := range files {
		moveToOutput(rootpath, x+".pdf")
		//deleteEpub(root, x+".epub")
	}
	log.Info("all done!!!")
}

func handleReader(reader *bufio.Reader) {
	printOutput := log.GetLevel() == log.DebugLevel
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if printOutput {
			fmt.Print(str)
		}
	}
}
