// #############################################################################
// # File: epub_to_pdf.go                                                      #
// # Project: convert                                                          #
// # Created Date: 2023/09/11 07:45:50                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 08:07:13                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package convert

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/TwiN/go-color"
	gopool "github.com/realjf/gopool/v2"
	commonUtils "github.com/realjf/utils"
	log "github.com/sirupsen/logrus"

	"github.com/realjf/epub2pdf/app/backend/model"
	"github.com/realjf/epub2pdf/app/backend/utils"
)

func EpubToPDF(ctx context.Context, req model.EpubToPDFReq) {
	files := utils.GetPaths(req.InputPath)
	if len(files) < req.JobsNum {
		req.JobsNum = len(files)
	}

	convertPool := gopool.NewPool(req.JobsNum)
	convertPool.SetTaskNum(len(files))
	if req.Timeout > 0 {
		convertPool.SetTimeout(req.Timeout)
	}
	convertPool.SetDebug(true)
	log.Info(color.InRed(strconv.Itoa(len(files))) + " files to be converted")

	// add task
	go func() {
		for _, filename := range files {
			x := filename
			myTaskFunc := func(param interface{}) (r interface{}, err error) {
				r, ok := param.(*model.FileObj)
				if !ok {
					return r, errors.New("task parameter is not file object type")
				}
				input_file := x.Abs()
				output_file := x.ToRootPath(req.OutputPath).ToAbs()
				log.Infof("ready to convert %s to %s ...\n", input_file, output_file)

				args := []string{input_file, output_file}
				cmd := commonUtils.NewCmd().SetDebug(true)
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
				input_file := param.(*model.FileObj).Abs()
				if req.IsDelete {
					err = utils.DeleteFile(input_file)
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
