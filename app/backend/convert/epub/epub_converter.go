// #############################################################################
// # File: epub_converter.go                                                   #
// # Project: epub                                                             #
// # Created Date: 2023/09/11 16:20:52                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 16:25:21                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"context"
	"os"

	"github.com/realjf/gopool/v2"
	commonUtils "github.com/realjf/utils"
	log "github.com/sirupsen/logrus"

	"github.com/realjf/epub2pdf/app/backend/model"
	"github.com/realjf/epub2pdf/app/backend/utils"
)

type EpubConverter interface {
	ToPDF(req *model.ConvertReq)
}

type epubConverter struct {
	ctx context.Context
}

func NewEpubConverter(ctx context.Context) EpubConverter {
	return &epubConverter{
		ctx: ctx,
	}
}

func (e *epubConverter) ToPDF(req *model.ConvertReq) {
	files := req.InputFiles
	if len(files) < req.JobsNum {
		req.JobsNum = len(files)
	}

	convertPool := gopool.NewPool(req.JobsNum)
	convertPool.SetTaskNum(len(files))
	if req.Timeout > 0 {
		convertPool.SetTimeout(req.Timeout)
	}
	convertPool.SetDebug(true)
	log.Debugf("%d files to be converted", len(files))

	// add task
	go func() {
		for _, filename := range files {
			x := filename
			err := convertPool.AddTask(func() {
				err := e.toPDF(x, req)
				if err != nil {
					log.Errorf("error converting: %v", err)
				}
			})
			if err != nil {
				panic("add task error")
			}
			log.Debugf("add file %s", x.FileName())
		}
	}()

	convertPool.Run()
	log.Debug("tasks is completed!!!")

	log.Infof("total: %d", convertPool.GetDoneNum())
	log.Infof("success: %d", convertPool.GetSuccessNum())
	log.Infof("fail: %d", convertPool.GetFailNum())
	log.Info("all done!!!")
}

func (e *epubConverter) toPDF(fileObj *model.FileObj, req *model.ConvertReq) (err error) {

	input_file := fileObj.Abs()
	output_file := fileObj.ToRootPath(req.OutputPath).ToAbs()
	log.Debugf("ready to convert %s to %s ...\n", input_file, output_file)

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
		log.Errorf("======== failed to convert %s ========\n%v", input_file, err)
		return
	} else {
		log.Infof("======== convert %s successfully ========", input_file)
	}

	if req.IsDelete {
		err = utils.DeleteFile(input_file)
		if err != nil {
			log.Errorf("========= delete %s error ========\n%v", input_file, err)
		} else {
			log.Infof("========= delete %s successfully ========", input_file)
		}
	}
	return
}
