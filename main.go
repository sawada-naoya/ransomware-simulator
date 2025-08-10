package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/sawada-naoya/ransomware-simulator/ui"
)


func main() {
    // ① アプリケーション作成
    a := app.New()
    // ② 新しいウィンドウ作成
    w := ui.CreateMainWindow(a)
    // ④ ウィンドウ表示 & イベントループ開始
    w.ShowAndRun()
}