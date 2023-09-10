// #############################################################################
// # File: app.go                                                              #
// # Project: frontend                                                         #
// # Created Date: 2023/09/10 23:02:14                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/10 23:43:43                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package frontend

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	a.win.SetContent(a.content)
	a.win.ShowAndRun()
}
