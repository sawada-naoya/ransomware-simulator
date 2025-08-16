package crypto

import (
	"os"
	"path/filepath"
)

func FakeEncryptFile(targetDir string) error {
	// 指定ディレクトリ配下の全ファイル・ディレクトリを再帰的に走査
  // path : 現在処理中のファイルパス
  // info : os.FileInfo でファイルの情報（サイズ、ディレクトリかどうか、更新時刻など）
  // err : 走査中に発生したエラー
	return filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error{
		if err != nil || info.IsDir() {
			return nil
		}
		// ファイルの拡張子を取得（例: "file.txt" → ".txt")
		ext := filepath.Ext(path)
		// .txt / .csv / .docx のみを対象とする（その他は何もしない）
		if ext == ".txt" || ext == ".csv" || ext == ".docx" {
			newName := path + ".locked"
			return os.Rename(path, newName)
		}
		return nil
	})
}

func FakeEncryptPath(path string) error {
	ext := filepath.Ext(path)
	switch ext {
		case ".txt", ".csv", ".docx":
			return os.Rename(path, path + ".locked")
		default:
			return nil
	}
}