package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(a fyne.App) fyne.Window {
    w := a.NewWindow("Ransomware Simulator")

    label := widget.NewLabel("あなたのファイルはロックされました")
    btn := widget.NewButton("支払い", func() {
        ShowPaymentDialog(w)
    })

    w.SetContent(container.NewVBox(
        label,
        btn,
    ))

    return w
}