package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)


func ShowPaymentDialog(parent fyne.Window) {
    dialog.ShowCustom("支払い完了", "閉じる", widget.NewLabel("復号キーが送信されました（という体験です）"), parent)
}