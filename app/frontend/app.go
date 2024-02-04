// #############################################################################
// # File: app.go                                                              #
// # Project: frontend                                                         #
// # Created Date: 2023/09/10 23:02:14                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/02/04 15:52:33                                        #
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
)

type FApp interface {
	SetContent(content fyne.CanvasObject)
	Run()
}

type fApp struct {
	app     fyne.App
	win     fyne.Window
	content fyne.CanvasObject
}

func NewApp(title string) FApp {
	a := &fApp{
		app: app.New(),
	}

	a.win = a.app.NewWindow(title)

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
	a.win.SetContent(a.content)
	a.win.ShowAndRun()
}
