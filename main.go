package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/sawada-naoya/ransomware-simulator/ui"
)


func main() {
    // アプリケーション作成
    a := app.New()
    // 新しいウィンドウ作成
    w := ui.CreateMainWindow(a)
    // ウィンドウ表示 & イベントループ開始
    w.ShowAndRun()
}