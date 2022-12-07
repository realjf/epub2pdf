package cmd

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
	"github.com/realjf/cgroup"
	"github.com/realjf/gopool"
	"github.com/realjf/relimit"
	"github.com/realjf/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	// convertCmd.Flags().StringVarP(&EbookConvertPath, "ebook-convert-path", "p", "ebook-convert", "The ebook-convert executable path")
	convertCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output.")
	convertCmd.Flags().BoolVarP(&MoreVerbose, "more-verbose", "m", false, "Output the converted output.")
	convertCmd.Flags().IntVarP(&JobsNum, "jobs", "j", 5, "Allow N jobs at once; infinite jobs with no arg.")
	convertCmd.Flags().StringVarP(&OutputPath, "output-path", "o", "", "Output path, by default, is the source directory.")
	convertCmd.Flags().BoolVarP(&Recursive, "recursive", "r", false, "Recursively search the directory that contains an epub file.")
	convertCmd.Flags().StringVarP(&ToBeConvertedPath, "path-to-convert", "f", "", "The path to be converted, required.")
	convertCmd.Flags().BoolVarP(&DeleteSource, "delete-source", "d", false, "Delete the source file when convert successfully.")
	convertCmd.Flags().IntVarP(&Timeout, "timeout", "t", 0, "Per-job's timeout value; the default is 10; and the unit of time is seconds.")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert -f=</path/to/epub_directory>",
	Short: "convert epub to pdf",
	Long:  `Convert the specified directory epub file to pdf file`,
	Run: func(cmd *cobra.Command, args []string) {
		// check ebook-convert exist
		if _, err := exec.LookPath(EbookConvertPath); err != nil {
			log.Fatal("The ebook-convert is not in your PAHT environment variable")
			os.Exit(1)
		}

		if ToBeConvertedPath == "" {
			log.Fatal("The path to be converted is empty, please confirm your path")
			os.Exit(1)
		}
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		// start to convert
		Convert()
	},
}

func getPaths(root string) []*FileObj {
	formats := map[string]bool{FILE_EXT_EPUB: true}

	files := []*FileObj{}
	err := filepath.Walk(root,
		func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if root == ".." {
				return nil
			}

			var rootpath string
			if root == "." {
				rootpath, err = filepath.Abs(ToBeConvertedPath)
				if err != nil {
					return err
				}
				rootpath = filepath.Join(rootpath, filepath.Dir(fp))
				if MoreVerbose {
					log.Info("Current directory1: " + rootpath)
				}
			} else {
				if info.IsDir() {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						return err
					}
					if MoreVerbose {
						log.Info("Current directory2: " + rootpath)
					}
				} else {
					rootpath, err = filepath.Abs(fp)
					if err != nil {
						return err
					}
					rootpath = filepath.Dir(rootpath)
					if MoreVerbose {
						log.Info("Current directory3: " + rootpath)
					}
				}

			}
			if !Recursive {
				if ro, err := filepath.Abs(root); err != nil {
					return err
				} else {
					if rootpath != ro {
						// Non recursive
						if MoreVerbose {
							log.Info("Non recursive: " + rootpath + "," + ro)
						}
						return nil
					}
				}
			}

			if !info.IsDir() && filepath.Ext(fp) != "" && formats[filepath.Ext(fp)] {
				fileObj := NewFileObj(fileNameWithoutExtension(info.Name()), filepath.Ext(fp), rootpath, FILE_EXT_PDF)
				if !fileExists(fileObj.Abs()) {
					log.Warn(color.InYellow("File[" + fileObj.Abs() + "] not found"))
					return nil
				}
				files = append(files, fileObj)
				if Verbose {
					log.Info("The path[" + fileObj.Abs() + "] to be converted")
				}
				return nil
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	// return strings
	return files

}

func Convert() {
	files := getPaths(ToBeConvertedPath)
	if len(files) < JobsNum {
		JobsNum = len(files)
	}

	convertPool := gopool.NewPool(JobsNum)
	convertPool.SetTaskNum(len(files))
	if Timeout > 0 {
		convertPool.SetTimeout(time.Second * time.Duration(Timeout))
	}
	convertPool.SetDebug(true)
	log.Info(color.InRed(strconv.Itoa(len(files))) + " files to be converted")

	relimiter := relimit.MustNewRelimit(80, 12*cgroup.Gigabyte, true)
	// relimiter.SetDebug(true)
	// relimiter.GetCmd().SetDebug(true)
	defer relimiter.Close()
	err := relimiter.StartByPid(os.Getpid())
	if err != nil {
		log.Error(err)
		return
	}

	// add task
	go func() {
		for _, filename := range files {
			x := filename
			myTaskFunc := func(param interface{}) (r interface{}, err error) {
				r, ok := param.(*FileObj)
				if !ok {
					return r, errors.New("task parameter is not file object type")
				}
				input_file := x.Abs()
				output_file := x.ToRootPath(OutputPath).ToAbs()
				log.Infof("ready to convert %s to %s ...\n", input_file, output_file)

				args := []string{input_file, output_file}
				cmd := utils.NewCmd().SetDebug(true)
				if os.Getenv("SUDO_USER") != "" {
					cmd.SetUsername(os.Getenv("SUDO_USER"))
					cmd.SetNoSetGroups(true)
				}
				envs := os.Environ()
				// envslices := []string{}
				// envs = append(envs, envslices...)
				cmd.SetEnv(envs)
				_, err = cmd.RunCommand("ebook-convert", args...)
				if err != nil {
					log.Error(color.InRed("======== failed to convert " + input_file + " ========"))
					log.Error(color.InRed(err.Error()))
					return
				} else {
					log.Info(color.InGreen("======== convert " + input_file + " successfully ========"))
				}

				return
			}
			myTaskCallbackFunc := func(param interface{}) (r interface{}, err error) {
				input_file := param.(*FileObj).Abs()
				if DeleteSource {
					err = deleteFile(input_file)
					if err != nil {
						log.Error("========= delete " + input_file + " error: " + color.InRed(err.Error()) + " ========")
					} else {
						log.Info(color.InRed("========= delete " + input_file + " successfully ========"))
					}
				}
				return err, nil
			}
			task := gopool.NewTask(myTaskFunc, myTaskCallbackFunc, x)
			err := convertPool.AddTask(task)
			if err != nil {
				panic("add task error")
			}
			log.Info("add task:" + x.FileName())
		}
	}()

	convertPool.Run()
	log.Info(color.InGreen("tasks is completed!!!"))

	log.Info(color.InGreen("total:" + strconv.Itoa(convertPool.GetDoneNum())))
	log.Info(color.InGreen("success:" + strconv.Itoa(convertPool.GetSuccessNum())))
	log.Info(color.InYellow("fail:" + strconv.Itoa(convertPool.GetFailNum())))
	log.Info(color.InRed("all done!!!"))
}
