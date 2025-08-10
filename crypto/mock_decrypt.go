package crypto

import (
	"os"
	"path/filepath"
	"strings"
)

func FakeDecryptFile(targetDir string) error {
	return filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error{
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".locked"){
			originalName := strings.TrimSuffix(path, ".locked")
			return os.Rename(path, originalName)
		}
		return nil
	})
}