package getFile

import (
	"os"
	"path/filepath"
)

// Get the last file inputed and return the file
func GetFile(dir string) (string, error) {
	var file string

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	entry := dirEntries[0]
	fileInfo, err := entry.Info()
	if err != nil {
		return "", err
	}
	file = filepath.Join(dir, fileInfo.Name())

	if file == "" {
		return "", os.ErrNotExist
	}

	return file, nil
}
