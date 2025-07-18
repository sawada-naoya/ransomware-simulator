package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LaunchUI() {
	a := app.New()
	w := a.NewWindow("⚠️ 警告 ⚠️")
	w.SetContent(container.NewVBox(
		widget.NewLabel("あなたのファイルはロックされました"),
	))
	w.ShowAndRun()
}