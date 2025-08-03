package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/sawada-naoya/ransomware-simulator/ui"
)


func main() {
    a := app.New()
    w := ui.CreateMainWindow(a)
    w.ShowAndRun()
}