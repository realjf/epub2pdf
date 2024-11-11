// #############################################################################
// # File: app.go                                                              #
// # Project: frontend                                                         #
// # Created Date: 2023/09/10 23:02:14                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/11 13:16:21                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package frontend

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/realjf/epub2pdf/app/config"
)

type FApp interface {
	SetContent(content fyne.CanvasObject)
	Run()
	IsShutdown() <-chan bool
}

type fApp struct {
	app       fyne.App
	win       fyne.Window
	content   fyne.CanvasObject
	closeChan chan bool
}

func NewApp(title string) FApp {
	a := &fApp{
		app:       app.New(),
		closeChan: make(chan bool, 1),
	}

	a.win = a.app.NewWindow(title)
	a.win.SetOnClosed(func() {
		a.closeChan <- true
	})

	runtime.SetFinalizer(a, closeFApp)

	return a
}

func closeFApp(a *fApp) {
	a.win.Close()
}

func (a *fApp) SetContent(content fyne.CanvasObject) {
	a.content = content
}

func (a *fApp) Run() {
	if a.content == nil {
		a.content = widget.NewLabel("Hello World!")
	}
	a.win.CenterOnScreen()
	// 设置窗口大小
	a.win.Resize(fyne.NewSize(config.GlobalConfig.Frontend.Screen.Width, config.GlobalConfig.Frontend.Screen.Height))
	a.win.SetContent(a.content)
	a.win.ShowAndRun()

}

func (a *fApp) IsShutdown() <-chan bool {
	return a.closeChan
}
