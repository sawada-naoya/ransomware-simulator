package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	mockcrypto "github.com/sawada-naoya/ransomware-simulator/crypto"
)

func CreateMainWindow(a fyne.App) fyne.Window {
    w := a.NewWindow("Ransomware Simulator")

    targetPath := desktopPath("dammy.txt")
    lockedPath := targetPath + ".locked"

    title := widget.NewLabel("あなたのファイルはロックされました")
	title.Alignment = fyne.TextAlignCenter

	center := container.NewVBox(
		title,
		widget.NewSeparator(),
	)

    canClose := false
    w.SetCloseIntercept(func() {
		if canClose {
			w.Close()
			return
		}
		dialog.ShowInformation(
			"送金待ち",
			"送金が完了するまでこのウィンドウは閉じられません。",
			w,
		)
	})

	w.SetFixedSize(true)
    w.SetFullScreen(true)

    btnDecrypt := widget.NewButton("送金", func() {
		if _, err := os.Stat(lockedPath); err != nil {
			dialog.ShowInformation("対象なし", "Desktop に dammy.txt.locked が見つかりません。", w)
			return
		}
		if err := mockcrypto.FakeDecryptPath(lockedPath); err != nil {
			dialog.ShowError(fmt.Errorf("復号失敗: %w", err), w)
			return
		}
        canClose = true
	})

    btnDecrypt.Importance = widget.LowImportance
	bottomRight := container.NewGridWrap(fyne.NewSize(120, 36), btnDecrypt)
	bottomBar := container.NewHBox(
		layout.NewSpacer(), // 左にスペーサーで右寄せ
		bottomRight,
	)

    w.SetContent(container.NewBorder(nil, bottomBar, nil, nil, center))

    go func() {
		if err := runEncryptionFlow(targetPath); err != nil {
			// 初期化系の致命的エラーだけはダイアログでも通知
			dialog.ShowError(err, w)
			return
		}
	}()

    w.Resize(fyne.NewSize(520, 220))
	w.CenterOnScreen()
	return w
}

func runEncryptionFlow(targetPath string) error {
	if err := ensureDummyText(targetPath); err != nil {
		return fmt.Errorf("ファイル作成失敗: %w", err)
	}
	if err := mockcrypto.FakeEncryptPath(targetPath); err != nil {
		return fmt.Errorf("暗号化失敗: %w", err)
	}
	return nil
}

func desktopPath(filename string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Desktop", filename)
}

func ensureDummyText(path string) error {
	// すでに存在するならそのまま
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	// なければ作る（安全なダミー内容）
	content := []byte("これはテストです。.\n")
	return os.WriteFile(path, content, 0o644)
}